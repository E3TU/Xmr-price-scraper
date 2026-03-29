package main

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
)

func main() {
	var url string = "https://www.tradingview.com/symbols/XMREUR/"

	browser := rod.New().MustConnect()
	defer browser.MustClose()

	page := browser.MustPage(url)

	time.Sleep(1000 * time.Millisecond)

	var text string

	for {

		el := page.MustElement(".last-zoF9r75I.js-symbol-last")

		text = el.MustText()

		if text == "" {
			fmt.Println("The string is empty")

			time.Sleep(1000 * time.Millisecond)

		} else {
			fmt.Println("Price of XMR:", text)
			break
		}
	}
}
