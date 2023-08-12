package main

import (
	"coreutils/cmd"
	"flag"
	"fmt"
	"os"
)

func main() {
	// var cmdString []string
	// cmdString = []string{"-n"}
	// cmd.Uname(cmdString)
	// cmd.Arch()

	m()

	// cmd.Base64("helloworld today is a good day")
	// cmd.Base64("aGVsbG93aGF0aGVmdWNrLHNka2Zqa2p3ZWYueWVhaCEhIQ==")
	// cmd.Base64("YWE=")
	//cmd.Base64("YWFh")
}

func m() {
	var CommandLine = flag.NewFlagSet(os.Args[1], flag.ExitOnError)
	if err := CommandLine.Parse(os.Args[1:]); err != nil {
		fmt.Println("error is ", err)
	}

	args := CommandLine.Args()
	// pgrm := strings.ToLower(args[0])
	pgrm := args[0]
	params := args[1:]

	switch pgrm {
	case "arch":
		cmd.Arch()
	case "uname":
		cmd.Uname(params)
	case "base64":
		cmd.Base64(params)
	case "basename":
		cmd.Basename(params)
	case "cat":
		cmd.Cat(params)
	default:
		_, err := fmt.Fprintf(os.Stderr, "no such program: %s\n", pgrm)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
