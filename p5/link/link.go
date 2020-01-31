package link

import (
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func ParseLink(url string) []string {
	var ret = []string{}

	r, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	doc, err := html.Parse(r.Body)
	if err != nil {
		log.Println(err)
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		// if n.Type == html.ElementNode && n.Data == "a" {
		// 	fmt.Println(n.Data)
		// }
		if n.Type == html.ElementNode && len(n.Attr) > 0 && n.Data == "a" {
			for _, v := range n.Attr {
				if v.Key == "href" && detDm(url, v.Val) {
					ret = append(ret, v.Val)
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return ret
}

func detDm(d string, sub string) bool {
	do := strings.TrimSpace(d)
	subd := strings.TrimSpace(sub)
	flst := strings.Split(subd, ".")
	fext := flst[len(flst)-1]

	if subd == "/" {
		return false
	} else if strings.Contains(do, subd) && do != subd && subd != "/" {
		return true
	} else if fext != "html" && strings.Contains(subd, ".") {
		return false
	} else if strings.Contains(subd, "https://") {
		if strings.Split(do, "/")[1] != strings.Split(subd, "/")[1] {
			return false
		}
	}
	return true
}
