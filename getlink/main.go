package main

import (
	"fmt"
	"getlink/link"
)

func main() {
	// fmt.Println(link.ParseToken("ex2.html"))
	s := link.ParseLink("ex3.html")
	for k, v := range s {
		fmt.Printf("key: %v, value: %v\n", k, v)
	}

}
