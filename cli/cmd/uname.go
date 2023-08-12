package cmd

import (
	"bytes"
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

// type empty struct{}

// uname only accept 1 flag

func Uname(s []string) {

	if len(s) > 1 {
		fmt.Fprintln(os.Stderr, "Too many flags")
		return
	}

	if len(s) == 0 {
		fmt.Fprintln(os.Stderr, "Please enter a flag")
		return
	}
	unixName := unix.Utsname{}
	if err := unix.Uname(&unixName); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	sysName := bytes.Trim(unixName.Sysname[:], "\x00")
	nodeName := bytes.Trim(unixName.Nodename[:], "\x00")
	release := bytes.Trim(unixName.Release[:], "\x00")
	version := bytes.Trim(unixName.Version[:], "\x00")
	machineName := bytes.Trim(unixName.Machine[:], "\x00")

	flags := map[string][]byte{
		"-a": bytes.Join([][]byte{sysName, nodeName, version, release, machineName}, []byte("\n")),
		"-s": sysName,
		"-n": nodeName,
		"-v": version,
		"-r": release,
		"-m": machineName,
	}

	for _, v := range s {
		if _, ok := flags[v]; !(ok) {
			fmt.Fprintln(os.Stderr, "flag not exists: ", v)
			return
		}

	}

	fmt.Fprintf(os.Stdout, "%s%s", flags[s[0]], "\n")
}
