package douban

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func testParse() {
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
