package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

type PriceResponse struct {
	Result map[string]struct {
		C []string `json:"c"`
	} `json:"result"`
	Error []string `json:"error"`
}

func GetMoneroPrice(c *gin.Context) {
	var url string = "https://api.kraken.com/0/public/Ticker?pair=XMREUR"

	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch data from kraken"})
		return
	}

	defer resp.Body.Close()

	var data PriceResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse Kraken response"})
		return
	}

	if len(data.Error) > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": data.Error})
		return
	}

	var price string
	for _, v := range data.Result {
		if len(v.C) > 0 {
			price = v.C[0]
		}
	}

	if price == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "price not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"price": price})
}

func main() {

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/price", GetMoneroPrice)

	r.Run(":8080")
}
