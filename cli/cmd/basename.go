package cmd

import (
	"fmt"
	"os"
	"strings"
)

func Basename(s []string) {
	if len(s) > 1 {
		fmt.Fprintln(os.Stderr, "too many inputs, please only use 1 input")
		return
	}

	if len(s) == 0 {
		fmt.Fprintln(os.Stderr, "Please enter a basename")
		return
	}

	fp := s[0]

	parts := strings.Split(fp, "/")
	fmt.Fprintln(os.Stdout, parts[len(parts)-1])
}
