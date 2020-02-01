package main

import (
	"fmt"
	"p5/link"
)

func main() {
	ret := link.ParseLink("https://huajie06.github.io")
	// ret := link.ParseLink("https://www.huajiezhang.com")
	for i, v := range ret {
		fmt.Println(i, v)
	}
}
