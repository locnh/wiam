package main

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// return only IP, for no-ip, ddns services alike
	r.GET("/ip", func(c *gin.Context) {
		if c.GetHeader("X-Forwarded-For") != "" {
			c.String(200, "%s", c.GetHeader("X-Forwarded-For"))
		} else {
			c.String(200, "%s", c.ClientIP())
		}
	})

	// return header, for debugging behind CDN
	r.GET("/header", func(c *gin.Context) {
		c.JSON(200, c.Request.Header)
	})

	r.GET("/request", func(c *gin.Context) {
		json, err := json.MarshalIndent(getRequestInfo(c), "", "  ")
		if err != nil {
			c.JSON(500, gin.H{"result": "failed"})
		} else {
			c.String(200, "%s", json)
		}
	})

	r.NoRoute(func(c *gin.Context) {
		json, err := json.MarshalIndent(getClientInfo(c), "", "  ")
		if err != nil {
			c.JSON(500, gin.H{"result": "failed"})
		} else {
			c.String(200, "%s", json)
		}
	})

	r.Run()
}

func getClientInfo(c *gin.Context) map[string]string {
	client := make(map[string]string)

	// Client IP
	if c.GetHeader("X-Forwarded-For") != "" {
		client["ip"] = c.GetHeader("X-Forwarded-For")
	} else {
		client["ip"] = c.ClientIP()
	}

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
