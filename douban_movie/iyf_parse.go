package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type MovieInfo struct {
	Region, Language, AddTime, DNRate string
}

type Iyf struct {
	Ret           int    `json:"ret"`
	Data          Data   `json:"data"`
	Msg           string `json:"msg"`
	IsSpecialArea int    `json:"isSpecialArea"`
}
type Result struct {
	AtypeName      string      `json:"atypeName"`
	VideoClassID   string      `json:"videoClassID"`
	Image          string      `json:"image"`
	Key            string      `json:"key"`
	Lang           string      `json:"lang"`
	Cid            string      `json:"cid"`
	LastName       string      `json:"lastName"`
	IsShowTodayNum bool        `json:"isShowTodayNum"`
	Title          string      `json:"title"`
	Hot            int         `json:"hot"`
	Rating         string      `json:"rating"`
	Year           int         `json:"year"`
	Regional       string      `json:"regional"`
	AddTime        string      `json:"addTime"`
	Directed       string      `json:"directed"`
	Starring       string      `json:"starring"`
	ShareCount     int         `json:"shareCount"`
	Dd             int         `json:"dd"`
	Dc             int         `json:"dc"`
	Comments       int         `json:"comments"`
	FavoriteCount  int         `json:"favoriteCount"`
	Contxt         string      `json:"contxt"`
	IsSerial       bool        `json:"isSerial"`
	Updateweekly   string      `json:"updateweekly"`
	CidMapper      string      `json:"cidMapper"`
	LastKey        string      `json:"lastKey"`
	Recommended    bool        `json:"recommended"`
	Updates        int         `json:"updates"`
	Tags           interface{} `json:"tags"`
	IsFilm         bool        `json:"isFilm"`
	IsDocumentry   bool        `json:"isDocumentry"`
	Labels         string      `json:"labels"`
	Charge         int         `json:"charge"`
}
type Info struct {
	Recordcount int      `json:"recordcount"`
	Result      []Result `json:"result"`
}
type Data struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Info []Info `json:"info"`
}

func main() {
	f, err := os.Open("./iyf.json")
	if err != nil {
		log.Fatal(err)
	}

	movieList := map[string]MovieInfo{}

	var ret Iyf
	if jerr := json.NewDecoder(f).Decode(&ret); jerr != nil {
		log.Fatal(jerr)
	}

	info := ret.Data.Info
	result := info[0].Result
	for _, v := range result {

		_, found := movieList[v.Title]
		if found == false {
			movieList[v.Title] = MovieInfo{v.Regional, v.Lang, v.AddTime, v.Rating}
		}
	}

	for key, value := range movieList {
		fmt.Printf("key: %v, value: %v\n", key, value)
	}

}
