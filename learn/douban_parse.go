package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type doubanIndividualMovie struct {
	Name            string
	Url             string
	DatePublished   string          `json:datePublished`
	Genre           []string        `json:genre`
	Duration        string          `json:duration`
	AggregateRating AggregateRating `json:aggregateRating`
	QueryDateTime   string
}

type AggregateRating struct {
	Type        string `json:"@type"`
	RatingCount string `json:"ratingCount"`
	BestRating  string `json:"bestRating"`
	WorstRating string `json:"worstRating"`
	RatingValue string `json:"ratingValue"`
}

func doubanReturnMoreUrl(s string) []byte {
	// will return "" if non-match
	re := regexp.MustCompile(`sid:\s(\d*)`)
	return re.Find([]byte(s))
}

func dbSearch(qs string) string {
	// TODO: if no results
	// build URL
	u, err := url.Parse("https://www.douban.com/search?")
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	q.Set("source", "suggest")
	q.Set("q", qs)
	u.RawQuery = q.Encode()
	//fmt.Println(u.String())

	// create a client
	var client http.Client
	req, err := http.NewRequest("GET", u.String(), nil)

	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")

	// send request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// parse response
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}

	var sid string
	doc.Find(".result .content").Each(func(i int, s *goquery.Selection) {
		onclick, _ := s.Find("a").Attr("onclick")
		if contains := strings.Contains(onclick, "dou_search_movie"); contains == true {
			sid += string(doubanReturnMoreUrl(onclick)) + "\n"
		}
	})

	if len(strings.Split(sid, "\n")) == 0 {
		return ""
	}
	return strings.Split(sid, "\n")[0]
}

func doubanGetIndMovieLocal() doubanIndividualMovie {
	f, err := os.Open("./douban_individual.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}

	var sBlock string
	doc.Find("script[type=\"application/ld+json\"]").Each(func(i int, s *goquery.Selection) {
		sBlock = s.Text()
	})

	var movieResult doubanIndividualMovie

	err = json.Unmarshal([]byte(sBlock), &movieResult)
	if err != nil {
		log.Fatal(err)
	}
	return movieResult
}

func doubanGetIndMovie(movieID string) doubanIndividualMovie {
	dbMovieUrl := fmt.Sprintf("https://movie.douban.com/subject/%v/", movieID)
	u, err := url.Parse(dbMovieUrl)
	if err != nil {
		log.Fatal(err)
	}

	var client http.Client
	req, err := http.NewRequest("GET", u.String(), nil)

	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var sBlock string
	// instead of getting the type=application/ld+json
	// there's a tag `div id="info"`, which can be a better solution
	doc.Find("script[type=\"application/ld+json\"]").Each(func(i int, s *goquery.Selection) {
		//doc.Find("script[type=application/ld+json]").Each(func(i int, s *goquery.Selection) {
		sBlock = s.Text()
	})

	//fmt.Println(sBlock)

	var movieResult doubanIndividualMovie

	err = json.Unmarshal([]byte(sBlock), &movieResult)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(movieResult)
	//fmt.Printf("movie title: %v\nlink: https://movie.douban.com%v\nRating count: %v\nAverage rating: %v", movieResult.Name, movieResult.Url, movieResult.AggregateRating.RatingCount, movieResult.AggregateRating.RatingValue)
	return movieResult
}

func doubanWrapper(s string) {
	// s := "双重躯体"
	//text := "我要回高三"
	// no result
	var movieInfo doubanIndividualMovie

	r := dbSearch(s)
	if r == "" {
		fmt.Println("no movie found")
	} else {
		movieId := strings.Fields(r)[1]
		fmt.Println(r)
		fmt.Println(movieId)

		//movieId := "35594791"
		movieInfo = doubanGetIndMovie(movieId)
	}
	timenow := time.Now().Format("2006-01-02 15:04:05")
	movieInfo.QueryDateTime = timenow
	fmt.Println(movieInfo)

}

func writeToJson() {
	// get one movie
	movieInfo := doubanGetIndMovieLocal()
	timenow := time.Now().Format("2006-01-02 15:04:05")
	movieInfo.QueryDateTime = timenow

	var movieJson []doubanIndividualMovie

	// append to movie slice
	movieJson = append(movieJson, movieInfo)

	movieJson = append(movieJson, movieInfo)

	b, err := json.MarshalIndent(movieJson, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("ify_movie_base.json", b, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// logic is
	// 1. get ify list of new movies
	// 2. compare it with existing local database
	// 3. if there's new one, run douban API
	// 4. append the results to the json
	writeToJson()

}
