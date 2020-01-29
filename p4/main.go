package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func parseHTML(s string) string {
	var ret []byte
	f, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}

	tk := html.NewTokenizer(f)
	for {
		n := tk.Next()
		if n == html.ErrorToken {
			break
		}

		if t := tk.Text(); byteValid(t) {
			ret = append(ret, t...)
			ret = append(ret, 32)
		}
	}
	return string(ret)
}

func main() {
	// fmt.Println(parseHTML("ex2.html"))
	parseHTML1("ex3.html")
}

func parseHTML1(s string) {
	fl, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := html.Parse(fl)
	if err != nil {
		log.Fatal(err)
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "div" {
			for _, a := range n.Attr {
				if a.Key == "class" {
					fmt.Println(a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)
}

func byteValid(b []byte) bool {
	var d = map[uint8]int{}
	for _, v := range b {
		if _, ok := d[v]; ok {
			d[v] = d[v] + 1
		} else {
			d[v] = 1
		}
	}
	// return d
	if d[10]+d[32] == len(b) || len(d) == 0 {
		return false
	}
	return true
}

func byteValid1(b []byte) map[uint8]int {
	var d = map[uint8]int{}
	for _, v := range b {
		if _, ok := d[v]; ok {
			d[v] = d[v] + 1
		} else {
			d[v] = 1
		}
	}
	return d
}
