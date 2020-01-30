package link

import (
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func ParseLink(s string) map[string]string {
	fl, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := html.Parse(fl)
	if err != nil {
		log.Fatal(err)
	}
	var str string = ""
	var mp = make(map[string]string)

	var f func(*html.Node)
	f = func(n *html.Node) {

		if n.Type == html.TextNode {
			// str += strings.Join(strings.Fields(n.Data), " ") + "=="
			str += strings.TrimSpace(n.Data) + " "
		}

		for _, v := range n.Attr {
			if v.Key == "href" { // this only returns href
				mp[v.Key] = v.Val
			}
		}
		mp["Text"] = strings.Join(strings.Fields(str), " ")
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}

	}
	f(doc)
	return mp
}

func ParseToken(s string) string {
	var ret []byte
	f, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}

	tk := html.NewTokenizer(f)
	for {
		n := tk.Next()
		if n == html.ErrorToken {
			break
		}

		if t := tk.Text(); byteValid(t) {
			ret = append(ret, t...)
			ret = append(ret, 32)
		}
	}
	return string(ret)
}

func byteValid(b []byte) bool {
	var d = map[uint8]int{}
	for _, v := range b {
		if _, ok := d[v]; ok {
			d[v] = d[v] + 1
		} else {
			d[v] = 1
		}
	}
	// return d
	if d[10]+d[32] == len(b) || len(d) == 0 {
		return false
	}
	return true
}
