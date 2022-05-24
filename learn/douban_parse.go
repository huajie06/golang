package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("./douban.html")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	ioR := bufio.NewReader(f)

	for {
		line, _, err := ioR.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(line)
	}
}
