package cdphz

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func TryClick() {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		// chromedp.WithDebugf(log.Printf),
		chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	var title, res string
	resp, err := chromedp.RunResponse(ctx,
		chromedp.Navigate(`https://ify.tv`),
		chromedp.Title(&title),
		chromedp.Text("body", &res, chromedp.NodeVisible),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Status)

	log.Printf("%s", title)

	fmt.Println(res)
}
