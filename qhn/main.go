package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

const apiBase string = "https://hacker-news.firebaseio.com/v0/"

type story struct {
	Source string
	By     string `json:"by"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Tt     string `json:"type"`
	Url    string `json:"url"`
}

type storyHandle struct {
	tmpl *template.Template
	s    []story
}

func (s storyHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := s.tmpl.Execute(w, s.s)

	if err != nil {
		http.Error(w, "page not found", http.StatusNotFound)
	}

}

func (s storyHandle) debug() {
	s.tmpl.Execute(os.Stdout, s.s)
}

func main() {
	// ids := getTopStory(5)
	// fmt.Println(ids)

	// var ids = []int{22378679, 22380380, 22378555, 22376794, 22380364, 22379969}
	// ret := returnIds(ids)
	// fmt.Println(ret)

	//dat := []story{{Url: "abc", Title: "yeah!!!"}, {Url: "abc", Title: "yeah!!!"}}

	ids := getTopStory(20)
	dat := returnIds(ids)

	tmpl := template.Must(template.ParseFiles("template.html"))
	h := storyHandle{tmpl, dat}

	m := http.NewServeMux()
	m.Handle("/", h)
	http.ListenAndServe(":8000", m)

	// h.debug()
}

func getTopStory(numOfItems int) []int {
	var err error

	url := fmt.Sprintf("%stopstories.json", apiBase)
	r, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	var ids []int
	b := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println(err)
	}

	err = b.Decode(&ids)
	if err != nil {
		log.Println(err)
	}
	return ids[:numOfItems]
}

func returnIds(ids []int) []story {
	ret := []story{}
	for _, v := range ids {
		ret = append(ret, getById(v))
	}
	return ret
}

func getById(id int) story {
	//url := apiBase + "item/" + strconv.Itoa(id) + ".json"
	url := fmt.Sprintf("%sitem/%d.json", apiBase, id)

	r, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		log.Println(err)
	}

	s1 := story{}
	err = json.Unmarshal(b, &s1)
	if err != nil {
		log.Println(err)
	}

	if len(strings.Split(s1.Url, "/")) >= 2 {
		s1.Source = strings.Split(s1.Url, "/")[2]
	}

	return s1

	//fmt.Println(s1)
}
