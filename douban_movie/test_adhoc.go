package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	u, err := url.Parse("https://movie.douban.com/subject/35594791/")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(u)
}
