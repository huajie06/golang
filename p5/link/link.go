package link

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

// ParseLink func... return []string
func ParseLink(url string) []string {
	var ret = []string{}
	r, err := http.Get(url)
	// fmt.Println("---", r.Request.URL, "---")

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
		if n.Type == html.ElementNode && len(n.Attr) > 0 && n.Data == "a" {
			for _, v := range n.Attr {
				if v.Key == "href" && detDm(r.Request.URL, v.Val) {
					ret = append(ret, v.Val)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	ret = dedupSlice(ret)

	for i, v := range ret {
		if strings.HasPrefix(v, "/") {
			ret[i] = fmt.Sprintf("%s://%s%s", r.Request.URL.Scheme, r.Request.URL.Host, v)
		}
	}

	return ret
}

func detDm(domain *url.URL, subdomain string) bool {
	d := domain
	sub := strings.TrimRight(subdomain, "/")

	flst := strings.Split(sub, ".")
	fext := flst[len(flst)-1]

	if strings.TrimRight(domain.String(), "/") == sub || sub == "/" {
		return false // this actually is fine? depends on the goal
	} else if strings.HasPrefix(sub, "/") &&
		(!(strings.Contains(sub, ".")) || (strings.Contains(sub, ".") && strings.ToLower(fext) == "html")) {
		return true
	} else if strings.HasPrefix(sub, "http") {
		u, _ := url.Parse(sub)
		if u.Scheme == d.Scheme && u.Host == d.Host {
			return true
		}
		return false
	} else {
		return false
	}
}

func dedupSlice(s []string) []string {
	var m = map[string]int{}
	var ret = []string{}
	for _, v := range s {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}

	for k, _ := range m {
		ret = append(ret, k)
	}
	return ret
}

var parsedLink = []string{}
var counter int = 0

// LoopPage func
func LoopPage(url []string, depth int) []string {

	var l = []string{}

	if counter >= depth {
		return parsedLink
	}

	for _, v := range url {
		if !(sliceContains(v, parsedLink)) {
			l = append(l, ParseLink(v)...)
			parsedLink = append(parsedLink, v)
		}
	}

	counter++

	return LoopPage(l, depth)
}

func sliceContains(s string, sl []string) bool {
	for _, v := range sl {
		if s == v {
			return true
		}
	}
	return false
}
