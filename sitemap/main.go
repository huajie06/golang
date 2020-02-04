package main

import (
	"flag"
	"fmt"
	"sitemap/link"
)

func main() {
	// ret := link.ParseLink("https://huajie06.github.io")
	// ret := link.ParseLink("https://huajie06.github.io")

	urlFlag := flag.String("url", "https://huajie06.github.io", "the url that you wants to parse")
	maxDepth := flag.Int("depth", 4, "the max depth to pages wants to parse")
	flag.Parse()

	pages := []string{*urlFlag}
	ret := link.LoopPage(pages, *maxDepth)
	for i, v := range ret {
		fmt.Println(i, v)
	}
}
