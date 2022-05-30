package main

import (
	"encoding/json"
	"errors"
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

// some TODO:
// change douban parser to look for `div id=info`, or fix the json with invalid chars
// duonao URL builder, to set the parameter in order
// movie database change name

const localDBname = "./ify_movie_base.json"
const localNewMovieList = "./newMovie.json"
const localDuonao = "./iyf.json"

// =========================below is from duonao=========================

type duonaoMovieInfoRet struct {
	Title, Region, Language, AddTime, DNRate string
}

type duonaoMovieInfo struct {
	Ret           int    `json:"ret"`
	Data          Data   `json:"data"`
	Msg           string `json:"msg"`
	IsSpecialArea int    `json:"isSpecialArea"`
}
type Result struct {
	// the key and lastkey, one of them can be used to create link
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

// =========================above is from duonao=========================

type doubanIndividualMovie struct {
	SearchedTitle   string
	ReturnReason    string
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

func contains(s []string, str string) bool {
	// function to check slice(need to be strings) contains an element
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func loadJsonDb() []doubanIndividualMovie {
	// read the database and return a list of movies already in DB
	var movieList []doubanIndividualMovie
	f, err := os.Open(localDBname)
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(b, &movieList)
	if err != nil {
		log.Fatal(err)
	}

	return movieList
}

func getDuonaoMovieList() []duonaoMovieInfoRet {

	absUrl := "https://m10.iyf.tv/api/list/Search?cinema=1&page=1&size=36&orderby=0&desc=1&cid=0,1,3&isserial=-1&isIndex=-1&isfree=-1&vv=9a939766a82c1047ef6da69eb23f62d5&pub=CJOrCpGnC3WnE2umEJ0sNrLJNp8sC34wCJWpEZWrDsOwEJWnDpewCJ0pOLyPip8o6J6SCnmmd1oQ6PepC9cmiZ2oCfip6Xenc3CRCYzDJbcPcCmP3OvOZ8tDZOpCZXZD6HXPM4uOZbbCsOtDp6"
	u, err := url.Parse(absUrl)

	//u, err := url.Parse("https://m10.iyf.tv/api/list/Search?")
	if err != nil {
		fmt.Println(errors.New("duonao api fail"))
		return []duonaoMovieInfoRet{}
	}
	// q := u.Query()
	// q.Set("cinema", "1")
	// q.Set("page", "1")
	// q.Set("size", "36")
	// q.Set("orderby", "0")
	// q.Set("desc", "1")
	// q.Set("cid", "0,1,3")
	// q.Set("isserial", "-1")
	// q.Set("isindex", "-1")
	// q.Set("isfree", "-1")
	// q.Set("vv", "9a939766a82c1047ef6da69eb23f62d5")
	// q.Set("pub", "CJOrCpGnC3WnE2umEJ0sNrLJNp8sC34wCJWpEZWrDsOwEJWnDpewCJ0pOLyPip8o6J6SCnmmd1oQ6PepC9cmiZ2oCfip6Xenc3CRCYzDJbcPcCmP3OvOZ8tDZOpCZXZD6HXPM4uOZbbCsOtDp6")

	//u.RawQuery = q.Encode()

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
	// text, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(text))
	var dnMovieList duonaoMovieInfo
	if err = json.NewDecoder(resp.Body).Decode(&dnMovieList); err != nil {
		log.Fatal(err)
	}

	var dnMovieListRet []duonaoMovieInfoRet
	info := dnMovieList.Data.Info
	result := info[0].Result
	for _, v := range result {
		dnMovieListRet = append(dnMovieListRet, duonaoMovieInfoRet{v.Title, v.Regional, v.Lang, v.AddTime, v.Rating})

	}

	return dnMovieListRet
}

func compareSrc(db []doubanIndividualMovie, dn []duonaoMovieInfoRet) []string {
	var dbTitleList, dnTitleList, returnList []string

	for _, dbv := range db {
		dbTitleList = append(dbTitleList, dbv.SearchedTitle)
	}

	for _, dnv := range dn {
		dnTitleList = append(dnTitleList, dnv.Title)
	}

	// check if all duonao list exist in douban database
	for _, v := range dnTitleList {
		if !contains(dbTitleList, v) {
			returnList = append(returnList, strings.TrimSpace(v))
		}
	}

	return returnList
}

func doubanReturnMoreUrl(s string) []byte {
	// will return "" if non-match
	re := regexp.MustCompile(`sid:\s(\d*)`)
	return re.Find([]byte(s))
}

func dbSearch(qs string) (string, error) {
	fmt.Printf("==========%v=========\n", qs)
	u, err := url.Parse("https://www.douban.com/search?")
	if err != nil {
		return "", err
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
		return "", nil
	}
	return strings.Split(sid, "\n")[0], nil
}

func doubanGetIndMovie(qs, movieID string) (doubanIndividualMovie, error) {
	fmt.Printf("==========%v=========\n", movieID)
	fmt.Printf("==========%v=========\n", qs)

	if qs == "" || movieID == "" {
		return doubanIndividualMovie{}, errors.New("query return none")
	}
	dbMovieUrl := fmt.Sprintf("https://movie.douban.com/subject/%v/", movieID)
	u, err := url.Parse(dbMovieUrl)
	if err != nil {
		return doubanIndividualMovie{SearchedTitle: qs, Url: movieID, ReturnReason: "url fail"}, err
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
		return doubanIndividualMovie{SearchedTitle: qs, Url: movieID, ReturnReason: "bad GET"}, err
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return doubanIndividualMovie{SearchedTitle: qs, Url: movieID, ReturnReason: "goquery error"}, err
	}

	var sBlock string
	// instead of getting the type=application/ld+json
	// there's a tag `div id="info"`, which can be a better solution
	doc.Find("script[type=\"application/ld+json\"]").Each(func(i int, s *goquery.Selection) {
		//doc.Find("script[type=application/ld+json]").Each(func(i int, s *goquery.Selection) {
		sBlock = s.Text()
	})

	if json.Valid([]byte(sBlock)) == false {
		return doubanIndividualMovie{SearchedTitle: qs, Url: movieID, ReturnReason: "bad json"}, errors.New("Invalid: can not parse json")
	}

	var movieResult doubanIndividualMovie

	err = json.Unmarshal([]byte(sBlock), &movieResult)
	if err != nil {
		return doubanIndividualMovie{SearchedTitle: qs, ReturnReason: "Unmarshal error"}, err
	}
	//fmt.Println(movieResult)
	//fmt.Printf("movie title: %v\nlink: https://movie.douban.com%v\nRating count: %v\nAverage rating: %v", movieResult.Name, movieResult.Url, movieResult.AggregateRating.RatingCount, movieResult.AggregateRating.RatingValue)
	return movieResult, nil
}

func doubanWrapper(s string) (doubanIndividualMovie, error) {
	// s := "双重躯体"
	//text := "我要回高三"
	// no result
	var movieInfo doubanIndividualMovie

	r, err := dbSearch(s)
	//fmt.Println(r)
	if err != nil {
		fmt.Println("search error")
		return doubanIndividualMovie{SearchedTitle: s, ReturnReason: "search error"}, err
	} else if len(strings.Fields(r)) < 2 {
		fmt.Println("no movie found, return non sid key")
		return doubanIndividualMovie{SearchedTitle: s, ReturnReason: "not found"}, err
	} else {
		movieId := strings.Fields(r)[1]
		//fmt.Println(r)
		//fmt.Println(movieId)

		//movieId := "35594791"
		movieInfo, err = doubanGetIndMovie(s, movieId)
		if err != nil {
			return doubanIndividualMovie{SearchedTitle: s, ReturnReason: fmt.Sprintf("%v", err)}, err
		}
	}
	timenow := time.Now().Format("2006-01-02 15:04:05")
	movieInfo.QueryDateTime = timenow
	movieInfo.SearchedTitle = s
	movieInfo.ReturnReason = "Succesfull"

	return movieInfo, nil

}

func writeToJson(movieJson []doubanIndividualMovie, fileLoc string) {

	b, err := json.MarshalIndent(movieJson, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(fileLoc, b, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("write file: %v succesfull!\n", fileLoc)
}

func main() {
	var movieSearchResult []doubanIndividualMovie

	// get movie database and duonao movie list
	dbMovieList := loadJsonDb()
	dnMovieList := getDuonaoMovieList()

	// return the movie whenre it needs to get douban info
	moviesToSearch := compareSrc(dbMovieList, dnMovieList)
	moviesToSearch2 := moviesToSearch[len(moviesToSearch)-7 : len(moviesToSearch)-3]
	fmt.Println(moviesToSearch)

	// control only search a few movies
	fmt.Println(moviesToSearch2)

	// getting info from douban
	if len(moviesToSearch) > 0 {
		for i, v := range moviesToSearch2 {
			fmt.Printf("index:%v, values to serach: %v\n", i, v)
			dbr, err := doubanWrapper(v)
			if err != nil {
				fmt.Println(err)
			}
			movieSearchResult = append(movieSearchResult, dbr)
			time.Sleep(10 * time.Second)
		}

		// creating database for all movies and a new movie file
		movieDB := append(movieSearchResult, dbMovieList...)

		writeToJson(movieDB, localDBname)
		writeToJson(movieSearchResult, localNewMovieList)

	} else {
		fmt.Println("no movies to seach")
	}

}
