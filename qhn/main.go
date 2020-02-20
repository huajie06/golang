package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const apiBase string = "https://hacker-news.firebaseio.com/v0"

type story struct {
	By    string `json:"by"`
	Id    int    `json:"id"`
	Title string `json:"title"`
	Tt    string `json:"type"`
	Url   string `json:"url"`
	//Kids  []int  `json:"kids"`
}

func main() {
	var err error

	url2 := "https://hacker-news.firebaseio.com/v0/topstories.json"
	r, err := http.Get(url2)
	if err != nil {
		log.Println(err)
	}

	defer r.Body.Close()

	var ids []int

	dec := json.NewDecoder(r.Body)
	err = dec.Decode(&ids)
	if err != nil {
		log.Println(err)
	}

	for _, v := range ids {
		fmt.Println(v)
	}
}

func getById() {

	url := "https://hacker-news.firebaseio.com/v0/item/22371629.json"
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

	fmt.Println(s1)

}
