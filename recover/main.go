package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"strings"

	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/panic/", panicDemo)
	mux.HandleFunc("/panic-after/", panicAfterDemo)
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/debug/", souceCodeHandle)
	log.Fatal(http.ListenAndServe(":8000", recoverMux(mux, true)))

	// r := makeLink(ss)
	// fmt.Println(r)
}

func recoverMux(h http.Handler, dev bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Println(r)
				stack := debug.Stack()
				log.Println(string(stack))
				if !dev {
					http.Error(w, "something went wrong", http.StatusInternalServerError)
					return
				} else {
					fmt.Fprintf(w, "<h1>%v</h1><br><pre>%s</pre>", r, makeLink(string(stack)))
					return
				}
			}
		}()
		h.ServeHTTP(w, r)
	}
}

func panicDemo(w http.ResponseWriter, r *http.Request) {
	funcThatPanics()
}

func panicAfterDemo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello!</h1>")
	funcThatPanics()
}

func funcThatPanics() {
	panic("Oh no!")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello!</h1>")
}

func souceCodeHandle(w http.ResponseWriter, r *http.Request) {
	// use this format
	// http://localhost:8000/debug/?path=/Users/huajiezhang/go/src/recover/main.go
	v := r.FormValue("path")
	if v == "" {
		http.Error(w, "Please enter a file", http.StatusInternalServerError)
		return
	}

	b := bytes.NewBuffer(nil)

	f, err := os.Open(v)
	if err != nil {
		http.Error(w, "File is not found. Please enter a file", http.StatusInternalServerError)
		return
	}

	_, err = b.ReadFrom(f)
	if err != nil {
		http.Error(w, "File Error. Please enter a right file", http.StatusInternalServerError)
		return
	}
	// fmt.Fprintln(w, b.String())

	lexer := lexers.Get("go")
	iterator, err := lexer.Tokenise(nil, b.String())
	style := styles.Get("github")
	if style == nil {
		style = styles.Fallback
	}

	formatter := html.New(html.TabWidth(2), html.WithLineNumbers(true), html.LineNumbersInTable(true))
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<style>pre { font-size: 1.2em; }</style>")
	formatter.Format(w, style, iterator)

	return
	// err = quick.Highlight(w, b.String(), "go", "html", "monokailight")

	// if err != nil {
	// 	log.Println(err)
	// }

}

func makeLink(s string) string {
	lines := strings.Split(s, "\n")
	for li, line := range lines {
		if len(line) == 0 || line[0] != '\t' {
			continue
		}

		file := ""
		lineNum := ""
		for i, ch := range line {
			if ch == ':' {
				file = line[1:i]
				lineNum = line[i+1:]
				break
			}
		}
		// fmt.Println(file)
		lines[li] = "\t<a href=\"/debug/?path=" + file + "\">" + file + "</a>" + ":" + lineNum
	}

	// fmt.Println(lines)
	return strings.Join(lines, "\n")

}
