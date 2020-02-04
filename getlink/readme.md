### a couple of thoughts

- use slice of struct to represent data instead of a map
- html.node needs to loop through
    - one of the ways to get the data
    ```go
	var f func(*html.Node)
	f = func(n *html.Node) {

		if n.Type == html.TextNode {
            //do something
		}

		for _, v := range n.Attr {
			if v.Key == "href" { // this only returns href
				//do something
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}

	}
	f(doc)
    ```