package main

import (
	"bufio"
	"fmt"
	"os"
)

// PROMPT used as flag
const PROMPT = "..."

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Printf(PROMPT)
		fmt.Println(scanner.Text())
		if scanner.Text() == "quit()" {
			return
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
