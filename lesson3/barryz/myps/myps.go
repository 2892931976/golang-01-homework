package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Proc struct {
	Pid     int
	Cmdline string
}

func main() {
	ps, err := listAllProcs()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, p := range ps {
		fmt.Printf("%d\t%s\n", p.Pid, p.Cmdline)
	}

	/* output:
	809     /usr/bin/lsmd-d
	814     /bin/dbus-daemon--system--address=systemd:--nofork--nopidfile--systemd-activation
	833     svscan/service
	836     superviselog
	8594    /opt/haproxy/sbin/haproxy-f/opt/haproxy/etc/rabbitmq_pool.cfg
	865     /sbin/rngd-f
	868     /usr/lib64/erlang/erts-5.10.4/bin/epmd-daemon
	*/
}

func listAllProcs() (ps []*Proc, err error) {
	var dirs []string
	dirSpec := "/proc"
	// mac test
	// dirSpec := "./proc"
	dirs, err = listDirsUnderDir(dirSpec)
	if err != nil {
		return
	}

	size := len(dirs)
	if size == 0 {
		err = fmt.Errorf("no dir found under %s", dirSpec)
		return
	}

	for i := 0; i < size; i++ {
		pid, e := strconv.Atoi(dirs[i])
		if e != nil {
			continue
		}

		cmdlineFile := fmt.Sprintf("%s/%d/cmdline", dirSpec, pid)

		// read cmdline file content
		cmdCnt, e := ioutil.ReadFile(cmdlineFile)
		if e != nil {
			continue
		}

		// if file's content is empty, skip it
		if len(cmdCnt) == 0 {
			continue
		}

		p := &Proc{Pid: pid, Cmdline: string(cmdCnt)}
		ps = append(ps, p)
	}

	return
}

func listDirsUnderDir(dirPath string) ([]string, error) {
	fs, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return []string{}, err
	}

	size := len(fs)
	if size == 0 {
		return []string{}, nil
	}

	// make a slice,  len is 0, and the cap is the size of fs
	ret := make([]string, 0, size)
	for i := 0; i < size; i++ {
		if fs[i].IsDir() {
			name := fs[i].Name()
			if name != "." && name != ".." {
				ret = append(ret, name)
			}
		}
	}

	return ret, nil
}
