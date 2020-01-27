package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
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

	mux := CustomizedMux()
	maphandler := CustHandler(dat, tmpl, mux)
	// for k, v := range *dat {
	// 	createHTML(k, tmpl, v)
	// }

	fmt.Println("server starting on: 8000")
	http.ListenAndServe(":8000", maphandler)

}

func CustomizedMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homePage)
	return mux
}

func homePage(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/intro", http.StatusFound)
}

func CustHandler(dat *map[string]Chapter, tmpl *template.Template, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if v, ok := (*dat)[strings.ToLower(strings.Trim(r.URL.Path, "/"))]; ok {
			tmpl.Execute(w, v)
			return
		}
		fallback.ServeHTTP(w, r)
	}
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
