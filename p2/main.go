package main

import (
	"fmt"
	"net/http"
	// "p2/urlshort"
)

// S1 ...
type S1 struct{}

func (S1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "user handler not handle or handlefunc")
}

func main() {
	// urlshort.Test()
	mux := defaultMux()
	// mapHandler := urlshort.Test()

	mux.Handle("/api/", S1{})

	fmt.Println("server starting on: 8080")
	http.ListenAndServe(":8080", mux)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world!")
}
