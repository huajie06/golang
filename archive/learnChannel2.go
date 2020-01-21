package main

import "fmt"

func main() {
	fmt.Println("start main")
	c := make(chan string)
	go func() {
		c <- "1"
		fmt.Println("in the middle")
		c <- "2"
		c <- "3"
	}()

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

	fmt.Println("start main")
}

/*
output is
---------
start main
in the middle
1
2
3
start main
*/

/*
the go routine (go func()) will get blocked when executing c<-1 and wait for <-c
it pass the data to the receiver <-c but fmt.Println is faster than that pass
*/
