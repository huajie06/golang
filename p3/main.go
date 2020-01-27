package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

func parseJson(s string) *map[string]Chapter {
	f, err := ioutil.ReadFile(s)
	if err != nil {
		panic(err)
	}
	var JsonData = make(map[string]Chapter)
	if err := json.Unmarshal(f, &JsonData); err != nil {
		panic(err)
	}
	return &JsonData
}

func main() {
	dat := parseJson("cyoa.json")

	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		panic(err)
	}
	for k, v := range *dat {
		fmt.Println(k)
		createHTML(k, tmpl, v)
	}

	// tmpl.Execute(os.Stdout, dat.Intro)

}

func createHTML(fname string, tmpl *template.Template, dt Chapter) {
	f, err := os.Create(fname + ".html")
	if err != nil {
		log.Print(err)
		return
	}
	defer f.Close()

	tmpl.Execute(f, dt)
}
