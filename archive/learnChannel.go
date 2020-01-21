package main

import (
	"fmt"
	"time"
)

// if attempting to send to a channel, then itself will be blocked until there's a receiver
// similary, if attempting to receive, it will get blocked untill there's a sender

func main() {
	fmt.Println("start Main method")

	c := make(chan string)

	go func(c chan string) {
		time.Sleep(1 * time.Second)
		c <- "time's up"
	}(c)

	nt := time.NewTimer(2 * time.Second)
	// NewTimer fires to nt.C when expires

	select {
	case <-c:
		fmt.Println("go routine fist")
	case <-nt.C:
		fmt.Println("outer timer first")
	}

	fmt.Println("End Main method")
}
