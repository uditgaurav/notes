package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/containerd/cgroups"
	"github.com/moby/locker"
	"github.com/pkg/errors"
	"k8s.io/klog"
)

var (
	// Possible cgroup subsystems
	cgroupSubsys = []string{"cpu", "memory", "systemd", "net_cls",
		"net_prio", "freezer", "blkio", "perf_event", "devices",
		"cpuset", "cpuacct", "pids", "hugetlb"}
)

const (
	containerdProtocolPrefix        = "containerd://"
	DefaultProcPrefix               = "/proc"
	PidNS                    NsType = "pid"
	MountNS                  NsType = "mnt"
	pausePath                       = "/usr/local/bin/pause"
	nsexecPath                      = "/usr/local/bin/nsexec"
	// uts namespace is not supported yet
	// UtsNS   NsType = "uts"
	IpcNS NsType = "ipc"
	NetNS NsType = "net"
)

func main() {
	id := "d9d2081e298c"
	pid := 14064
	path := pidPath(int(pid))

	var out, stderr bytes.Buffer

	// id, err := FormatContainerID(cid)
	// if err != nil {
	// 	klog.Errorf("Err: %v", err)
	// }

	fmt.Println("*************:  " + id)

	cgroup, err := findValidCgroup(path, id)
	if err != nil {
		klog.Errorf("Err: %v", err)
		os.Exit(1)
	}

	control, err := cgroups.Load(cgroups.V1, cgroups.StaticPath(cgroup))
	if err != nil {
		klog.Infof("err: %v", err)
		os.Exit(1)
	}
	fmt.Println(&control)

	// processBuilder := DefaultProcessBuilder("stress-ng", "--cpu", "2").EnablePause()
	// if true {
	// 	processBuilder = processBuilder.SetNS(uint32(pid), PidNS)
	// 	fmt.Println("************************************************************")
	// 	fmt.Println(processBuilder)
	// 	fmt.Println("************************************************************")

	// }

	// cmd := processBuilder.Build()

	// err = StartProcess(cmd)
	// if err != nil {
	// 	klog.Infof("%v", err)
	// 	os.Exit(1)
	// }

	cmd := exec.Command("/bin/bash", "-c", "/usr/local/bin/pause /usr/local/bin/nsexec -p /proc/14064/ns/pid -- stress-ng --cpu 2")
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Start()
	if err != nil {
		klog.Infof(fmt.Sprint(err) + ": " + stderr.String())
	}
	klog.Infof("%v", out.String())
	time.Sleep(2 * time.Second)

	pid_stress := cmd.Process.Pid
	klog.Infof("Pid is: %d", pid)

	klog.Info("Start process successfully")

	// procState, err := process.NewProcess(int32(pid_stress))
	// if err != nil {
	// 	klog.Infof("%v", err)
	// 	os.Exit(1)
	// }
	// ct, err := procState.CreateTime()
	// if err != nil {
	// 	klog.Infof("%v", err)
	// 	os.Exit(1)
	// }
	// klog.Infof("Start Time: %v", ct)

	if err = control.Add(cgroups.Process{Pid: cmd.Process.Pid}); err != nil {
		klog.Infof("Errorrrrrrrrrrrrrrrrrrr: %v", err)
	}

	if err := cmd.Process.Signal(syscall.SIGCONT); err != nil {
		klog.Errorf("Error: %v", err)
	}

	klog.Infof("Process ID: %v", pid_stress)

	klog.Info("done")
}

/////////////////////////////////////////////////////////////////////////////

// DaemonServer represents a grpc server for tc daemon
type DaemonServer struct {
	crClient                 ContainerRuntimeInfoClient
	backgroundProcessManager BackgroundProcessManager

	IPSetLocker *locker.Locker
}

type ContainerRuntimeInfoClient interface {
	GetPidFromContainerID(containerID string) (uint32, error)
	ContainerKillByContainerID(ctx context.Context, containerID string) error
	FormatContainerID(ctx context.Context, containerID string) (string, error)
}

type ExecStressRequest struct {
	Scope                ExecStressRequest_Scope `protobuf:"varint,1,opt,name=scope,proto3,enum=pb.ExecStressRequest_Scope" json:"scope,omitempty"`
	Target               string                  `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	Stressors            string                  `protobuf:"bytes,3,opt,name=stressors,proto3" json:"stressors,omitempty"`
	EnterNS              bool                    `protobuf:"varint,4,opt,name=enterNS,proto3" json:"enterNS,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

type ExecStressRequest_Scope int32

// ProcessBuilder builds a exec.Cmd for daemon
type ProcessBuilder struct {
	cmd  string
	args []string

	nsOptions []nsOption

	pause    bool
	localMnt bool

	identifier *string

	ctx context.Context
}
type nsOption struct {
	Typ  NsType
	Path string
}

type NsType string

var nsArgMap = map[NsType]string{
	MountNS: "m",
	// uts namespace is not supported by nsexec yet
	// UtsNS:   "u",
	IpcNS: "i",
	NetNS: "n",
	PidNS: "p",
	// user namespace is not supported by nsexec yet
	// UserNS:  "U",
}

// ManagedProcess is a process which can be managed by backgroundProcessManager
type ManagedProcess struct {
	*exec.Cmd

	// If the identifier is not nil, process manager should make sure no other
	// process with this identifier is running when executing this command
	Identifier *string
}

// Build builds the process
func (b *ProcessBuilder) Build() *ManagedProcess {
	args := b.args
	cmd := b.cmd

	if len(b.nsOptions) > 0 {
		args = append([]string{"--", cmd}, args...)
		for _, option := range b.nsOptions {
			args = append([]string{"-" + nsArgMap[option.Typ], option.Path}, args...)
		}

		if b.localMnt {
			args = append([]string{"-l"}, args...)
		}
		cmd = nsexecPath
	}

	// if b.pause {
	// 	args = append([]string{cmd}, args...)
	// 	cmd = pausePath
	// }

	klog.Infof("build command", "command", cmd+" "+strings.Join(args, " "))

	command := exec.CommandContext(b.ctx, cmd, args...)
	command.SysProcAttr = &syscall.SysProcAttr{}
	command.SysProcAttr.Pdeathsig = syscall.SIGTERM

	return &ManagedProcess{
		Cmd:        command,
		Identifier: b.identifier,
	}
}

// // StartProcess manages a process in manager
// func StartProcess(cmd *ManagedProcess) error {
// 	var identifierLock *sync.Mutex
// 	if cmd.Identifier != nil {

// 		identifierLock.Lock()
// 	}

// 	err := cmd.Start()
// 	if err != nil {
// 		klog.Errorf("fail to start process", err)
// 		return err
// 	}

// 	_ = pid_stress
// 	procState, err := process.NewProcess(int32(pid_stress))
// 	if err != nil {
// 		return err
// 	}

// 	_, err = procState.CreateTime()
// 	if err != nil {
// 		return err
// 	}

// 	type ProcessPair struct {
// 		Pid        int
// 		CreateTime int64
// 	}

// 	// ProcessPair is an identifier for process
// 	go func() {
// 		err := cmd.Wait()
// 		if err != nil {
// 			err, ok := err.(*exec.ExitError)
// 			if ok {
// 				status := err.Sys().(syscall.WaitStatus)
// 				if status.Signaled() && status.Signal() == syscall.SIGTERM {
// 					klog.Info("process stopped with SIGTERM signal")
// 				}
// 			} else {
// 				klog.Errorf("process exited accidentally: %v", err)
// 			}
// 		}

// 		klog.Info("process stopped")

// 	}()

// 	return nil
// }

// BackgroundProcessManager manages all background processes
type BackgroundProcessManager struct {
	deathSig    *sync.Map
	identifiers *sync.Map
}

// NewBackgroundProcessManager creates a background process manager
func NewBackgroundProcessManager() BackgroundProcessManager {
	return BackgroundProcessManager{
		deathSig:    &sync.Map{},
		identifiers: &sync.Map{},
	}
}

// SetNS sets the namespace of the process
func (b *ProcessBuilder) SetNS(pid uint32, typ NsType) *ProcessBuilder {
	return b.SetNSOpt([]nsOption{{
		Typ:  typ,
		Path: GetNsPath(pid, typ),
	}})
}

// SetNSOpt sets the namespace of the process
func (b *ProcessBuilder) SetNSOpt(options []nsOption) *ProcessBuilder {
	b.nsOptions = append(b.nsOptions, options...)

	return b
}

// GetNsPath returns corresponding namespace path
func GetNsPath(pid uint32, typ NsType) string {
	return fmt.Sprintf("%s/%d/ns/%s", DefaultProcPrefix, pid, string(typ))
}

// DefaultProcessBuilder returns the default process builder
func DefaultProcessBuilder(cmd string, args ...string) *ProcessBuilder {
	return &ProcessBuilder{
		cmd:        cmd,
		args:       args,
		nsOptions:  []nsOption{},
		pause:      false,
		identifier: nil,
		ctx:        context.Background(),
	}
}

// EnablePause enables pause for process
func (b *ProcessBuilder) EnablePause() *ProcessBuilder {
	b.pause = true

	return b
}

// findValidCgroup ...
func findValidCgroup(path cgroups.Path, target string) (string, error) {
	for _, subsys := range cgroupSubsys {
		p, err := path(cgroups.Name(subsys))
		if err != nil {
			klog.Error(err, "Failed to retrieve the cgroup path", "subsystem", subsys, "target", target)
			continue
		}
		if strings.Contains(p, target) {
			return p, nil
		}
	}
	return "", fmt.Errorf("never found valid cgroup for %s", target)
}

func pidPath(pid int) cgroups.Path {
	p := fmt.Sprintf("/proc/%d/cgroup", pid)
	paths, err := parseCgroupFile(p)
	if err != nil {
		return errorPath(errors.Wrapf(err, "parse cgroup file %s", p))
	}
	return existingPath(paths, pid, "")
}

func parseCgroupFile(path string) (map[string]string, error) {
	f, err := os.Open(path)
	if err != nil {
		klog.Infof("Err: %v", err)
	}
	defer f.Close()
	return parseCgroupFromReader(f)
}

func parseCgroupFromReader(r io.Reader) (map[string]string, error) {
	var (
		cgroups = make(map[string]string)
		s       = bufio.NewScanner(r)
	)
	for s.Scan() {
		var (
			text  = s.Text()
			parts = strings.SplitN(text, ":", 3)
		)
		if len(parts) < 3 {
			return nil, fmt.Errorf("invalid cgroup entry: %q", text)
		}
		for _, subs := range strings.Split(parts[1], ",") {
			if subs != "" {
				cgroups[subs] = parts[2]
			}
		}
	}

	if err := s.Err(); err != nil {
		klog.Infof("Err: %v", err)
	}

	return cgroups, nil
}

func errorPath(err error) cgroups.Path {
	return func(_ cgroups.Name) (string, error) {
		return "", err
	}
}

func existingPath(paths map[string]string, pid int, suffix string) cgroups.Path {
	// localize the paths based on the root mount dest for nested cgroups
	for n, p := range paths {
		dest, err := getCgroupDestination(pid, string(n))
		if err != nil {
			return errorPath(err)
		}
		rel, err := filepath.Rel(dest, p)
		if err != nil {
			return errorPath(err)
		}
		if rel == "." {
			rel = dest
		}
		paths[n] = filepath.Join("/", rel)
	}
	return func(name cgroups.Name) (string, error) {
		root, ok := paths[string(name)]
		if !ok {
			if root, ok = paths[fmt.Sprintf("name=%s", name)]; !ok {
				return "", cgroups.ErrControllerNotActive
			}
		}
		if suffix != "" {
			return filepath.Join(root, suffix), nil
		}
		return root, nil
	}
}

func getCgroupDestination(pid int, subsystem string) (string, error) {
	// use the process's mount info
	p := fmt.Sprintf("/proc/%d/mountinfo", pid)
	f, err := os.Open(p)
	if err != nil {
		return "", err
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		fields := strings.Fields(s.Text())
		for _, opt := range strings.Split(fields[len(fields)-1], ",") {
			if opt == subsystem {
				return fields[3], nil
			}
		}
	}
	if err := s.Err(); err != nil {
		return "", err
	}
	return "", fmt.Errorf("never found desct for %s", subsystem)
}

// type ContainerdClientInterface interface {
// 	LoadContainer(ctx context.Context, id string) (containerd.Container, error)
// }

// type ContainerdClient struct {
// 	client ContainerdClientInterface
// }

// FormatContainerID strips protocol prefix from the container ID
func FormatContainerID(containerID string) (string, error) {
	if len(containerID) < len(containerdProtocolPrefix) {
		return "", fmt.Errorf("container id %s is not a containerd container id", containerID)
	}
	if containerID[0:len(containerdProtocolPrefix)] != containerdProtocolPrefix {
		return "", fmt.Errorf("expected %s but got %s", containerdProtocolPrefix, containerID[0:len(containerdProtocolPrefix)])
	}
	return containerID[len(containerdProtocolPrefix):], nil
}
