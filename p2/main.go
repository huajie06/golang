package main

import (
	"fmt"
	"net/http"
	// "p2/urlshort"
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

	// var pathtourls = map[string]string{
	// 	"/test": "https://www.google.come",
	// }

	mux := defaultMux()

	// mapHandler := urlshort.MapHandler()

	fmt.Println("server starting on: 8000")
	http.ListenAndServe(":8000", mux)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", S1{})
	mux.Handle("/newpage1", S2{"https://www.google.com", 301})
	return mux
}
