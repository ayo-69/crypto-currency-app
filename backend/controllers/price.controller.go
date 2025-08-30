package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

var priceCache = cache.New(5*time.Minute, 10*time.Minute)

func GetPrice(c *gin.Context) {
	symbols := strings.ReplaceAll(c.DefaultQuery("symbols", "bitcoin,ethereum"), " ", "")
	currencies := strings.ReplaceAll(c.DefaultQuery("currencies", "usd"), " ", "")

	cacheKey := symbols + "|" + currencies

	if cached, found := priceCache.Get(cacheKey); found {
		c.JSON(http.StatusOK, gin.H{
			"symbols":    symbols,
			"currencies": currencies,
			"prices":     cached,
			"cached":     true,
		})
		return
	}

	url := "https://api.coingecko.com/api/v3/simple/price?ids=" +
		strings.ToLower(symbols) +
		"&vs_currencies=" + strings.ToLower(currencies)

	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch prices"})
		return
	}
	defer resp.Body.Close()

	var data map[string]map[string]float64
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse response"})
		return
	}

	priceCache.Set(cacheKey, data, cache.DefaultExpiration)

	c.JSON(http.StatusOK, gin.H{
		"symbols":    symbols,
		"currencies": currencies,
		"prices":     data,
		"cached":     false,
	})
}
