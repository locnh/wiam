package main

import (
	bgo "github.com/digitalcrab/browscap_go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type clientInfo struct {
	ip          string
	browser     string
	os          string
	country     string
	countrycode string
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		client := getClientInfo(c)
		c.JSON(200, gin.H{
			"ip":          client.ip,
			"country":     client.country,
			"countrycode": client.countrycode,
			"browser":     client.browser,
			"os":          client.os,
		})
	})
	r.Run()
}

func getClientInfo(c *gin.Context) clientInfo {
	client := clientInfo{}
	var ok bool

	// Client IP
	if c.GetHeader("HTTP_CF_CONNECTING_IP") != "" {
		client.ip = c.GetHeader("HTTP_CF_CONNECTING_IP")
	} else {
		client.ip = c.ClientIP()
	}

	// GeoIP
	client.countrycode = c.GetHeader("HTTP_CF_IPCOUNTRY")
	if client.country, ok = countrycode[client.countrycode]; !ok {
		client.country = "unknown"
	}

	// Client OS, Browser, Country, Countrycode
	if err := bgo.InitBrowsCap("browscap.ini", false); err != nil {
		log.Error(err)
	}

	browser, ok := bgo.GetBrowser(c.GetHeader("User-Agent"))
	if !ok || browser == nil {
		client.os = "unknown"
		client.browser = "unknown"
	} else {
		client.os = browser.Platform
		client.browser = browser.Browser
	}

	return client
}
