package main

import (
	"fmt"
	"net/http"
	// "p2/urlshort"
)

func main() {
	// urlshort.Test()
	// mux := defaultMux()
	// mapHandler := urlshort.Test()

	h := hello
	http.HandleFunc("/hello", h)
	fmt.Println("server starting on: 8080")
	http.ListenAndServe(":8080", nil)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world!")
}
