package main

import (
	"fmt"
	"p5/link"
)

func main() {
	// ret := link.ParseLink("https://huajie06.github.io")
	// ret := link.ParseLink("https://huajie06.github.io")
	// pages := []string{"https://huajie06.github.io"}
	pages := []string{"https://www.huajiezhang.com"}
	ret := link.LoopPage(pages, 6)
	for i, v := range ret {
		fmt.Println(i, v)
	}

	// fmt.Println(link.ParseLink("https://www.huajiezhang.com/explore"))
}
