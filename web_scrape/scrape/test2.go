package cdphz

import (
	"context"
	"fmt"
	"log"

	"github.com/chromedp/chromedp"
)

func RunTest1() {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var res string
	resp, err := chromedp.RunResponse(ctx,
		chromedp.Navigate(`https://pkg.go.dev/time`),
		chromedp.Text(`.Documentation-overview`, &res, chromedp.NodeVisible),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
	fmt.Println("third status code:", resp.Status)
}
