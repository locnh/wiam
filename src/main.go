package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// return only IP, for no-ip, ddns services alike
	r.GET("/ip", func(c *gin.Context) {
		c.String(http.StatusOK, "%s", c.ClientIP())
	})

	// return headers
	r.GET("/headers", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, c.Request.Header)
	})

	// return cookies
	r.GET("/cookies", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, c.Request.Cookies())
	})

	// return User-Agent
	r.GET("/ua", func(c *gin.Context) {
		userAgent := c.GetHeader("X-Real-User-Agent")
		if userAgent == "" {
			userAgent = c.Request.UserAgent()
		}
		c.String(http.StatusOK, "%s", userAgent)
	})

	// return status code
	r.GET("/status/:code", func(c *gin.Context) {
		code := c.Param("code")
		// Convert string to integer
		var statusCode int
		_, err := fmt.Sscanf(code, "%d", &statusCode)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid status code")
			return
		}
		c.Status(statusCode)
	})

	// redirect n times
	r.GET("/redirect/:n", func(c *gin.Context) {
		n := c.Param("n")
		// Convert string to integer
		var redirectCount int
		_, err := fmt.Sscanf(n, "%d", &redirectCount)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid redirect count")
			return
		}

		// Limit the number of redirects to prevent infinite loops
		if redirectCount < 0 || redirectCount > 10 {
			c.String(http.StatusBadRequest, "Redirect count must be between 0 and 10")
			return
		}

		// If redirect count is 0, return 200 OK
		if redirectCount == 0 {
			c.String(http.StatusOK, "OK")
			return
		}

		// Perform the redirect
		// Redirect to the same endpoint with decremented count
		redirectURL := fmt.Sprintf("/redirect/%d", redirectCount-1)
		c.Redirect(http.StatusFound, redirectURL)
	})

	// basic authentication
	r.GET("/basic-auth/:username/:password", func(c *gin.Context) {
		username := c.Param("username")
		password := c.Param("password")

		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.Header("WWW-Authenticate", "Basic realm=\"Restricted Area\"")
			c.Status(http.StatusUnauthorized)
			return
		}

		// Check if the Authorization header starts with "Basic "
		if !strings.HasPrefix(auth, "Basic ") {
			c.Header("WWW-Authenticate", "Basic realm=\"Restricted Area\"")
			c.Status(http.StatusUnauthorized)
			return
		}

		// Extract the base64 encoded part
		encoded := strings.TrimPrefix(auth, "Basic ")
		if encoded == "" {
			c.Header("WWW-Authenticate", "Basic realm=\"Restricted Area\"")
			c.Status(http.StatusUnauthorized)
			return
		}

		// Decode the base64 encoded credentials
		decoded, err := base64.StdEncoding.DecodeString(encoded)
		if err != nil {
			c.Header("WWW-Authenticate", "Basic realm=\"Restricted Area\"")
			c.Status(http.StatusUnauthorized)
			return
		}

		// Split the decoded string into username and password
		creds := strings.Split(string(decoded), ":")
		if len(creds) != 2 {
			c.Header("WWW-Authenticate", "Basic realm=\"Restricted Area\"")
			c.Status(http.StatusUnauthorized)
			return
		}

		// Check if the credentials match
		if creds[0] == username && creds[1] == password {
			c.String(http.StatusOK, "Access granted")
		} else {
			c.Header("WWW-Authenticate", "Basic realm=\"Restricted Area\"")
			c.Status(http.StatusUnauthorized)
		}
	})

	// return all request details for all methods
	r.Any("/request", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, getAllRequestInfo(c))
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

func getAllRequestInfo(c *gin.Context) map[string]interface{} {
	result := make(map[string]interface{})

	// Client IP
	result["origin_ip"] = c.ClientIP()

	// Query strings
	query := make(map[string]string)
	for key, values := range c.Request.URL.Query() {
		if len(values) > 0 {
			query[key] = strings.Join(values, ", ")
		}
	}
	result["query"] = query

	// URL parameters
	params := make(map[string]string)
	for _, param := range c.Params {
		params[param.Key] = param.Value
	}
	result["params"] = params

	// Headers
	headers := make(map[string]string)
	for key, values := range c.Request.Header {
		if len(values) > 0 {
			headers[key] = strings.Join(values, ", ")
		}
	}
	result["headers"] = headers

	// Method
	result["method"] = c.Request.Method

	// Request URI
	result["uri"] = c.Request.RequestURI

	// Host
	result["host"] = c.Request.Host

	// User-Agent
	userAgent := c.GetHeader("X-Real-User-Agent")
	if userAgent == "" {
		userAgent = c.Request.UserAgent()
	}
	result["user_agent"] = userAgent

	// Payload (body)
	body := ""
	if c.Request.Body != nil {
		// For demonstration purposes, we'll read the body
		// In a real scenario, you might want to handle this differently
		bodyBytes, _ := c.GetRawData()
		body = string(bodyBytes)
	}
	result["payload"] = body

	// Cookies
	cookies := make(map[string]string)
	for _, cookie := range c.Request.Cookies() {
		cookies[cookie.Name] = cookie.Value
	}
	result["cookies"] = cookies

	return result
}
