package cdphz

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func Run1() {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		// chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	url := `https://www.iyf.tv/list?cid=0,1,3&orderBy=0&desc=true`
	// fmt.Println(url)

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// navigate to a page, wait for an element, click
	var res string
	resp, err := chromedp.RunResponse(ctx,
		// chromedp.Navigate(`https://pkg.go.dev/time`),
		chromedp.Navigate(url),
		// wait for footer element is visible (ie, page is loaded)
		chromedp.WaitVisible(`body > app-pager`),
		// find and click "Example" link
		chromedp.Text(`.search-results`, &res, chromedp.NodeVisible),
		// retrieve the text of the textarea
		// chromedp.Value(`#example-After textarea`, &example),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Status)
	fmt.Println(resp.Headers)
	fmt.Println(resp.StatusText)

	log.Printf("%s", res)
}
