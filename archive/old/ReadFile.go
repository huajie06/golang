package archive

import (
	"fmt"
	"os"
)

func mainRead() {
	f := createFile("data/test.txt")
	defer fileClose(f)
	fileWrite(f)
}

func createFile(p string) *os.File {
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func fileClose(f *os.File) {
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func fileWrite(f *os.File) {
	fmt.Fprintln(f, "bunch of data")
}
