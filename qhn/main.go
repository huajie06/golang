package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
	"sync"
	"syscall"
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

type templateData struct {
	Item []story
	Time time.Duration
}

func main() {
	tmpl := template.Must(template.ParseFiles("template.html"))
	h := topStoryHandle(tmpl, 30)
	h1 := newStoryHandle(tmpl, 30)

	m := http.NewServeMux()
	m.Handle("/", h)
	m.Handle("/hello", h1)

	fmt.Println("site started at port:8000.")
	if err := http.ListenAndServe(":8000", m); err != nil {
		log.Println(err)
	}

}

func debug() {
	ids, _ := getNewStory(30)
	dat := getIds(ids)
	fmt.Println(dat)
}

type cacheStory struct {
	numStory        int
	cache           []story
	cacheExpiration time.Time
	cacheMutext     sync.Mutex
}

func (cs *cacheStory) returnCacheIds() []story {
	cs.cacheMutext.Lock()
	defer cs.cacheMutext.Unlock()

	if time.Now().Sub(cs.cacheExpiration) < 0 {
		return cs.cache
	}
	ids, _ := getNewStory(cs.numStory)
	cs.cache = getIds(ids)
	cs.cacheExpiration = time.Now().Add(10 * time.Second)
	return cs.cache
}

func newStoryHandle(tmpl *template.Template, numStories int) http.HandlerFunc {

	cs := cacheStory{numStory: numStories}
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		dat := cs.returnCacheIds()
		s := templateData{dat, time.Now().Sub(start)}

		err := tmpl.Execute(w, s)
		if err != nil {
			if !(errors.Is(err, syscall.EPIPE)) {
				log.Println(err)
				http.Error(w, "page not found", http.StatusInternalServerError)
			}
			return
		}
	}
}

func topStoryHandle(tmpl *template.Template, numStories int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		dat := returnCacheIds(numStories)
		s := templateData{dat, time.Now().Sub(start)}

		err := tmpl.Execute(w, s)
		if err != nil {
			// 2020/02/23 22:21:45 http: superfluous
			// response.WriteHeader call from
			// main.hnHandle.func1 (main.go:73)
			// the cause is -> write tcp
			// [::1]:8000->[::1]:64515: write: broken pipe

			// filter out the borken pip error
			if !(errors.Is(err, syscall.EPIPE)) {
				log.Println(err)
				http.Error(w, "page not found", http.StatusInternalServerError)
			}
			return
		}
	}
}

func getNewStory(numOfItems int) ([]int, error) {
	var err error
	url := fmt.Sprintf("%snewstories.json", apiBase)
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

var (
	cacheMutext     sync.Mutex
	cacheExpiration time.Time
	cache           []story
)

func returnCacheIds(num int) []story {
	cacheMutext.Lock()
	defer cacheMutext.Unlock()

	if time.Now().Sub(cacheExpiration) < 0 {
		return cache
	}
	ids, _ := getTopStory(num)
	cache = getIds(ids)
	cacheExpiration = time.Now().Add(40 * time.Second)
	return cache
}

func getIds(ids []int) []story {
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
	targetUrl := fmt.Sprintf("%sitem/%d.json", apiBase, id)
	ret := story{}

	r, err := http.Get(targetUrl)
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
		ret.Source = strings.TrimPrefix(strings.Split(ret.Url, "/")[2], "www.")
	}

	return ret, nil
}
