package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
	"text/template"
	"time"
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
	Item []story
	Time time.Duration
}

func main() {
	h := hnHandle()

	m := http.NewServeMux()
	m.Handle("/", h)

	fmt.Println("site started at port:8000.")
	if err := http.ListenAndServe(":8000", m); err != nil {
		log.Println(err)
	}
}

func debug() {
	ids, err := getTopStory(5)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(ids)

	dat := returnIds(ids)
	fmt.Println(dat)
}

func hnHandle() http.HandlerFunc {

	tmpl := template.Must(template.ParseFiles("template.html"))

	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		ids, err := getTopStory(30)
		if err != nil {
			http.Error(w, "page not found", http.StatusNotFound)
		}

		dat := returnIds(ids)

		s := storyHandle{dat, time.Now().Sub(start)}

		err = tmpl.Execute(w, s)
		if err != nil {
			http.Error(w, "page not found", http.StatusNotFound)
		}
	}
}

func getTopStory(numOfItems int) ([]int, error) {
	var err error

	url := fmt.Sprintf("%stopstories.json", apiBase)
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	var ids []int
	b := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err != nil {
		return nil, err
	}

	err = b.Decode(&ids)
	if err != nil {
		return nil, err
	}
	return ids[:numOfItems], nil
}

func returnIds(ids []int) []story {
	type chResult struct {
		s   story
		ind int
		err error
	}

	storyCh := make(chan chResult)

	for i, v := range ids {
		go func(i, v int) {
			s, err := getById(v)
			if err != nil {
				storyCh <- chResult{err: err}
			}
			storyCh <- chResult{s: s, ind: i}
		}(i, v)
	}

	var results []chResult
	for i := 0; i < len(ids); i++ {
		results = append(results, <-storyCh)
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].ind < results[j].ind
	})

	var ret []story
	for i := 0; i < len(ids); i++ {
		if results[i].err != nil {
			continue
		}
		ret = append(ret, results[i].s)
	}
	return ret
}

func getById(id int) (story, error) {
	url := fmt.Sprintf("%sitem/%d.json", apiBase, id)
	ret := story{}

	r, err := http.Get(url)
	if err != nil {
		return ret, err
	}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(b, &ret)
	if err != nil {
		return ret, err
	}

	if len(strings.Split(ret.Url, "/")) >= 2 {
		ret.Source = strings.Split(ret.Url, "/")[2]
	}

	return ret, nil

	//fmt.Println(ret)
}
