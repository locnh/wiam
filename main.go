package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// return only IP, for no-ip, ddns services alike
	r.GET("/ip", func(c *gin.Context) {
		c.String(http.StatusOK, "%s", c.ClientIP())
	})

	// return header, for debugging behind CDN
	r.GET("/header", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, c.Request.Header)
	})

	r.GET("/request", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, getRequestInfo(c))
	})

	r.NoRoute(func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, getClientInfo(c))
	})

	r.Run()
}

func getClientInfo(c *gin.Context) map[string]string {
	client := make(map[string]string)

	// Client IP
	client["ip"] = c.ClientIP()

	// CloudFlare
	if c.GetHeader("cf-ipcountry") != "" {
		client["country"] = countrycode[c.GetHeader("cf-ipcountry")]
	}

	// CloudFront
	if c.GetHeader("Cloudfront-Viewer-Country-Name") != "" {
		client["country"] = c.GetHeader("Cloudfront-Viewer-Country-Name")
	}
	if c.GetHeader("Cloudfront-Viewer-City") != "" {
		client["city"] = c.GetHeader("Cloudfront-Viewer-City")
	}

	return client
}

func getRequestInfo(c *gin.Context) map[string]string {
	client := make(map[string]string)

	client["host"] = c.Request.Host
	client["uri"] = c.Request.RequestURI
	client["method"] = c.Request.Method

	if c.GetHeader("Cloudfront-Viewer-Http-Version") != "" {
		client["proto"] = c.GetHeader("Cloudfront-Viewer-Http-Version")
	}

	return client
}
