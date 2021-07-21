package main

import (
	"net/http"

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

	// return only IP, for no-ip, ddns services alike
	r.GET("/ip", func(c *gin.Context) {
		if c.GetHeader("HTTP_CF_CONNECTING_IP") != "" {
			c.String(http.StatusOK, "%s", c.GetHeader("HTTP_CF_CONNECTING_IP"))
		} else {
			c.String(http.StatusOK, "%s", c.ClientIP())
		}
	})

	// return header, for debugging behind CDN
	r.GET("/header", func(c *gin.Context) {
		c.JSON(200, c.Request.Header)
	})

	r.NoRoute(func(c *gin.Context) {
		client := getClientInfo(c)

		c.JSON(200, gin.H{
			"ip":          client.ip,
			"country":     client.country,
			"countrycode": client.countrycode,
			"browser":     client.browser,
			"os":          client.os,
			"req_host":    c.Request.Host,
			"req_uri":     c.Request.RequestURI,
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
	client.countrycode = c.GetHeader("cf-ipcountry")
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
