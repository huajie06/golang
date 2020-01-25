package main

import (
	"fmt"
	"net/http"
	"p2/urlshort"
)

// S1 ...
type S1 struct{}

func (S1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Use handler not handle or handlefunc")
}

// S2 ...
type S2 struct {
	url  string
	code int
}

func (s S2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, s.url, s.code)
}

func main() {

	var pathsToUrls = map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
		"/abc":            "https://www.google.com",
	}
	mux := defaultMux()

	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yaml := `
	- path: /urlshort
	  url: https://github.com/gophercises/urlshort
	- path: /urlshort-final
	  url: https://github.com/gophercises/urlshort/tree/solution
	`

	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("server starting on: 8000")
	http.ListenAndServe(":8000", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", S1{})
	mux.Handle("/newpage1", S2{"https://www.google.com", 301})
	mux.HandleFunc("/aa", hello)
	mux.HandleFunc("/bb", hello2)

	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func hello2(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://www.huajiezhang.com", 301)
}
