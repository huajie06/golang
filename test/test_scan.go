package main

import (
	"bufio"
	"fmt"
	"os"
)

// PROMPT used as flag
const PROMPT = "go>> "

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for fmt.Printf(PROMPT); scanner.Scan(); fmt.Printf(PROMPT) {
		ln := scanner.Text()
		fmt.Println(ln)
		if ln == "quit()" {
			return
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
