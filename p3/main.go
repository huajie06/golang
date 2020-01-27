package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

type JsonData struct {
	Intro struct {
		Title   string   `json:"title"`
		Story   []string `json:"story"`
		Options []struct {
			Text string `json:"text"`
			Arc  string `json:"arc"`
		} `json:"options"`
	} `json:"intro"`
	NewYork struct {
		Title   string   `json:"title"`
		Story   []string `json:"story"`
		Options []struct {
			Text string `json:"text"`
			Arc  string `json:"arc"`
		} `json:"options"`
	} `json:"new-york"`
	Debate struct {
		Title   string   `json:"title"`
		Story   []string `json:"story"`
		Options []struct {
			Text string `json:"text"`
			Arc  string `json:"arc"`
		} `json:"options"`
	} `json:"debate"`
	SeanKelly struct {
		Title   string   `json:"title"`
		Story   []string `json:"story"`
		Options []struct {
			Text string `json:"text"`
			Arc  string `json:"arc"`
		} `json:"options"`
	} `json:"sean-kelly"`
	MarkBates struct {
		Title   string   `json:"title"`
		Story   []string `json:"story"`
		Options []struct {
			Text string `json:"text"`
			Arc  string `json:"arc"`
		} `json:"options"`
	} `json:"mark-bates"`
	Denver struct {
		Title   string   `json:"title"`
		Story   []string `json:"story"`
		Options []struct {
			Text string `json:"text"`
			Arc  string `json:"arc"`
		} `json:"options"`
	} `json:"denver"`
	Home struct {
		Title   string        `json:"title"`
		Story   []string      `json:"story"`
		Options []interface{} `json:"options"`
	} `json:"home"`
}

func parseJson(s string) *JsonData {
	f, err := ioutil.ReadFile(s)
	if err != nil {
		panic(err)
	}
	var dat = JsonData{}
	if err := json.Unmarshal(f, &dat); err != nil {
		panic(err)
	}
	return &dat
}

func main() {
	dat := parseJson("cyoa.json")
	fmt.Println(dat.Intro.Title)
	// fmt.Println(dat.Intro.Story)
	// fmt.Println(dat.Intro.Options)
	// for _, v := range dat.Intro.Story {
	// 	fmt.Println(v)
	// }

	// for _, v := range dat.Intro.Options {
	// 	fmt.Println(v.Text, v.Arc)
	// }

	tmpl, err := template.ParseFiles("t.html")
	if err != nil {
		panic(err)
	}

	f, err := os.Create("home.html")
	if err != nil {
		log.Print(err)
		return
	}
	defer f.Close()

	// tmpl.Execute(os.Stdout, dat.Intro)
	tmpl.Execute(f, dat.Intro)

}
