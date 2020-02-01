package link

import (
	"log"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

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
	return dedupSlice(ret)
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
			m[v] += 1
		} else {
			m[v] = 1
		}
	}

	for k, _ := range m {
		ret = append(ret, k)
	}
	return ret
}
