package archive

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/html"
)

func mainHtmlPkg() {
	s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`

	_, err := html.Parse(strings.NewReader(s))
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(doc.Type)

	nt := html.NewTokenizer(strings.NewReader(s))

	for {
		z := nt.Next()
		if z == html.ErrorToken {
			//return nt.Err()
			break
		}
		//fmt.Println(string(nt.Text()))
		// output text with space

		//fmt.Println(nt.Token())
		// as raw html - token with text

		//b, y := nt.TagName()
		//fmt.Println(string(b), y)
		//tag: p, a and etc, y: has attr or not

		//b, y, x := nt.TagAttr()
		//fmt.Println(string(b), string(y), x)
		// x: y/n has attr, b is the tag, y is the value

		if t := nt.Text(); len(t) > 0 {
			fmt.Println(string(t))
		}
	}
}
