package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func main() {
	r := gin.Default()
	// Add auth
	r.Use(func(c *gin.Context) {
		if !strings.HasPrefix(c.Request.URL.Path, "/swagger/") {
			if c.Request.URL.Query().Get("key") != getEnv("API_KEY", "x") {
				c.AbortWithError(401, fmt.Errorf("unauthorized"))
				return
			}
			c.Header("cache-control", "no-cache")
		}
	})
	// Create router
	v1 := r.Group("/v1")
	{
		// Get a Kv value
		v1.GET("/kv/:kvKey", GetValue)
		// Add new Kv value
		v1.PUT("/kv/:kvKey", PutValue)
		// Delete a kv key
		v1.DELETE("/kv/:kvKey", DeleteValue)
	}
	// Add docs
	r.Run()
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}
