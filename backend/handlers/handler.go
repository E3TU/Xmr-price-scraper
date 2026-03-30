package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-rod/rod"
)

func GetPrice(c *gin.Context) {
	var url string = "https://www.tradingview.com/symbols/XMREUR/"
	var message string = "The string is empty..."

	browser := rod.New().MustConnect()
	defer browser.MustClose()

	page := browser.MustPage(url)

	time.Sleep(1000 * time.Millisecond)

	var text string

	for {

		el := page.MustElement(".last-zoF9r75I.js-symbol-last")

		text = el.MustText()

		if text == "" {
			fmt.Println(message)

			time.Sleep(1000 * time.Millisecond)

		} else {
			fmt.Println("Price of XMR:", text)
			c.JSON(http.StatusOK, gin.H{"Price of XMR": text})
			break
		}
	}
}
