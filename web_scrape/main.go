package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

const (
	cnbcURL string = "https://www.cnbc.com/id/20409666/device/rss/rss.html"
)

type cnbcXML struct {
	XMLName xml.Name `xml:"rss"`
	Items   []item   `xml:"channel>item"`
}

type item struct {
	Link  string `xml:"link"`
	Title string `xml:"title"`
}

func main() {
	var v cnbcXML

	req, err := http.NewRequest("GET", cnbcURL, nil)
	if err != nil {
		log.Println(err)
	}
	r, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer r.Body.Close()
	d := xml.NewDecoder(r.Body)

	d.Strict = false

	err = d.Decode(&v)
	if err != nil {
		log.Println(err)
	}

	for _, v := range v.Items {
		fmt.Println(v.Link, v.Title)
	}

}
