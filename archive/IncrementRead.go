package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, _ := os.Open("//Users/huajiezhang/go/src/project1/data/problems.csv") // For read access.
	defer file.Close()

	var b = make([]byte, 10)
	var offs int64 = 0

	for {
		p, err := file.ReadAt(b, offs)
		if err == io.EOF {
			fmt.Printf("%v", string(b[:p]))
			break
		}
		fmt.Printf("%v", string(b))
		offs = offs + 10

	}

}
