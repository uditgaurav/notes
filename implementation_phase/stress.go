package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/containerd/cgroups"
	"github.com/pkg/errors"
	"k8s.io/klog"
)

var (
	// Possible cgroup subsystems
	cgroupSubsys = []string{"cpu", "memory", "systemd", "net_cls",
		"net_prio", "freezer", "blkio", "perf_event", "devices",
		"cpuset", "cpuacct", "pids", "hugetlb"}
)

func main() {
	id := "6b518933400f"
	pid := 4071
	path := pidPath(int(pid))

	var out, stderr bytes.Buffer

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
	processes, _ := control.Processes(cgroups.Devices, false)
	for _, p := range processes {
		klog.Infof("Pid: %v", p.Pid)
	}

	cmd := exec.Command("/bin/bash", "-c", "/usr/local/bin/pause /usr/local/bin/nsutil -t 4071 -p -- stress-ng --cpu 2 --timeout 20s")
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Start()
	time.Sleep(5 * time.Second)

	go AbortWatcher(cmd)

	klog.Infof("Target Container Process: %v", pid)
	klog.Infof("cmd.ProcessState.Pid: %v", cmd.Process.Pid)
	if err != nil {
		klog.Infof(fmt.Sprint(err) + ": " + stderr.String())

	}
	klog.Infof(fmt.Sprint(err) + ": " + stderr.String())

	klog.Infof("start error cmd.ProcessState.ExitCode : %d", cmd.ProcessState.ExitCode)
	klog.Info("Start process successfully")

	if err = control.Add(cgroups.Process{Pid: cmd.Process.Pid}); err != nil {
		klog.Infof("Error: %v", err)
	}

	time.Sleep(5 * time.Second)

	processes, _ = control.Processes(cgroups.Devices, false)
	for _, p := range processes {
		klog.Infof("Pid: %v", p.Pid)
	}

	if err := cmd.Process.Signal(syscall.SIGCONT); err != nil {
		klog.Infof("Error in resuming the process: %v", err)
	}

	for {
		if err := cmd.Process.Signal(syscall.SIGCONT); err != nil {
			klog.Infof("Error in resuming the process: %v", err)
		}

		klog.Info("send signal to resume process")
		time.Sleep(2 * time.Second)

		comm, err := ReadCommName(cmd.Process.Pid)
		if err != nil {
			klog.Infof("%v", err)
		}
		if comm != "pause\n" {
			klog.Infof("pause has been resumed comm: %v", comm)
			break
		}
		klog.Infof("the process hasn't resumed, step into the following loop comm: ", comm)
	}

	klog.Infof("Process ID: %v", cmd.Process.Pid)
	klog.Info("done")
	cmd.Wait()
	err = TerminateProcess(cmd.Process.Pid)
	if err != nil {
		klog.Info("Process termination failed: %v",err)
	}
}

//TerminateProcess will terminate a given process
func TerminateProcess(pid int)error{
	p, err := os.FindProcess(int(pid))
	if err != nil {
		klog.Errorf("unreachable path. `os.FindProcess` will never return an error on unix",err)
		return err
	}
	err = p.Signal(syscall.SIGTERM)

	if err != nil && err.Error() != "os: process already finished" {
		klog.Errorf("error while killing process",err)
		return err
	}
	return nil
}

func AbortWatcher(cmd *exec.Cmd) {

	// signChan channel is used to transmit signal notifications.
	signChan := make(chan os.Signal, 1)
	// Catch and relay certain signal(s) to signChan channel.
	signal.Notify(signChan, os.Interrupt, syscall.SIGTERM)

loop:
	for {
		select {
		case <-signChan:
			err := cmd.Wait()
			if err != nil {
				err, ok := err.(*exec.ExitError)
				if ok {
					status := err.Sys().(syscall.WaitStatus)
					if status.Signaled() && status.Signal() == syscall.SIGTERM {
						klog.Info("process stopped with SIGTERM signal")
					}
				} else {
					klog.Infof("process exited accidentally", err)
				}
			}

			klog.Info("process stopped")

			break loop
		}
	}
	os.Exit(1)
}

// ReadCommName returns the command name of process
func ReadCommName(pid int) (string, error) {
	f, err := os.Open(fmt.Sprintf("/proc/" + strconv.Itoa(pid) + "/comm"))
	if err != nil {
		return "", err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}

	return string(b), nil
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
