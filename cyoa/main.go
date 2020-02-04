package main

import (
	"encoding/json"
	"fmt"
	"html/template"
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

type handler struct {
	Story *map[string]Chapter
	tmpl  *template.Template
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// process the url 1.default page. 2.remove "/" which might not be the optimal solution
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	p := strings.Trim(path, "/")

	if v, ok := (*h.Story)[p]; ok {
		h.tmpl.Execute(w, v)
		return
	}
	http.Error(w, "Chapter not found.", http.StatusNotFound)
}

func parseJson(s string) *map[string]Chapter {
	f, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}
	var JsonData = make(map[string]Chapter)
	if err := json.NewDecoder(f).Decode(&JsonData); err != nil {
		log.Fatal(err)
	}
	return &JsonData
}

func main() {
	dat := parseJson("cyoa.json")

	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		log.Fatal(err)
	}
	h := handler{dat, tmpl}

	// fmt.Println(h.Story)

	mux := CustomizedMux()
	// maphandler := CustHandler(dat, tmpl, mux)

	mux.Handle("/", h)

	fmt.Println("server starting on: 8000")
	http.ListenAndServe(":8000", mux)

}

func CustomizedMux() *http.ServeMux {
	mux := http.NewServeMux()
	// mux.HandleFunc("/", homePage)
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
