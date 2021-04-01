package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

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

	id := "6a9e2a28d86de12e67fa9c3af2eaddbbe03f6414b10a4238baed43f1e2816727"
	pid := 27179
	path := pidPath(int(pid))

	fmt.Println("*******************        path      ******************************************88")
	fmt.Println(path)
	fmt.Println("*************************        path      ************************************88")

	cgroup, err := findValidCgroup(path, id)
	if err != nil {
		klog.Errorf("Err: %v", err)
	}
	fmt.Println("*******************        cgroup      ******************************************88")
	fmt.Println(cgroup)
	fmt.Println("*************************        cgroup      ************************************88")

	// control, err := cgroups.Load(cgroups.V1, cgroups.StaticPath(""))
	// if err != nil {
	// 	klog.Infof("err: %v", err)
	// }

	// stats, err := control.Stat()
	// if err != nil {
	// 	klog.Infof("err: %v", err)
	// }
	// klog.Infof("stats: %v", stats)

	// if err := control.Add(cgroups.Process{Pid: 1234}); err != nil {
	// 	klog.Infof("err: %v", err)
	// }
}

/////////////////////////////////////////////////////////////////////////////
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
		return nil, err
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
		return nil, err
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
