package cmd

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

type empty struct{}

func Cat(s []string) {

	e := empty{}
	if len(s) == 0 {
		fmt.Fprintln(os.Stderr, "Please enter a file")
		return
	}

	fmt.Fprintln(os.Stdout, s)

	// -e = display $ at the end of each line
	// -n = display line number
	// -t = display tabs
	flags := map[string]empty{"-e": e, "-n": e, "-t": e}

	// if cat -d, meaning there's a `-`, it will try to parse the flag
	flag, fname := s[0], s[0]
	if strings.Contains(flag, "-") {
		if _, ok := flags[fname]; !ok {
			fmt.Fprintln(os.Stderr, "Flag is not supported.")
			return
		}
		fname = s[1]
	}

	f, err := os.Open(fname)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	var buf bytes.Buffer
	_, err = buf.ReadFrom(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	switch flag {
	case "-e":
		displayDollar(buf)
	case "-n":
		displayLineNum(buf)
	case "-t":
		displayTab(buf)
	}

	//fmt.Fprintf(os.Stdout, buf.String())
}

func displayDollar(b bytes.Buffer) {
	fmt.Fprintf(os.Stdout, "%s\n", bytes.ReplaceAll(b.Bytes(), []byte("\n"), []byte("$\n")))
	return
}

func displayLineNum(b bytes.Buffer) {
	lines := bytes.Split(b.Bytes(), []byte("\n"))
	padLen := len(lines)/10 + 2

	formatter := fmt.Sprintf("%%%dd  %%s\n", padLen)
	for i, v := range lines[:len(lines)-1] {
		fmt.Fprintf(os.Stdout, formatter, i+1, v)
	}
	return
}

func displayTab(b bytes.Buffer) {
	fmt.Fprintf(os.Stdout, "%s\n", bytes.ReplaceAll(b.Bytes(), []byte("\t"), []byte("^I")))
	return
}
