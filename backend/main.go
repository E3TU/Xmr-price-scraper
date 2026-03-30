package main

import (
	"xmr-price-scraper/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/price", handlers.GetPrice)

	r.Run(":8080")
}
