package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	// Request the HTML page.
	res, err := http.Get("http://metalsucks.net")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("article .post-title").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		title := s.Find("a").Text()
		fmt.Printf("Review %d: %s\n", i, title)
	})
}

func main() {

	// ExampleScrape()
	f, err := os.Open("./douban.html")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Println(err)
	}

	doc.Find(".result .content").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".subject-cast").Text()
		rating := s.Find(".rating_nums").Text()
		fmt.Printf("index:%d, title: %v, rating: %v\n", i, title, rating)
	})

	// ioR := bufio.NewReader(f)
	// for {
	// 	line, _, err := ioR.ReadLine()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(line)
	// }
}
