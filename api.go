package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func formatKvKey(key string) string {
	return "kv:" + key
}

// /v1/kv/:kvKey
const kvKeyParam = "kvKey"

func GetValue(c *gin.Context) {
	kvKey := formatKvKey(c.Param(kvKeyParam))
	kvValue, err := rdb.Get(kvKey).Result()
	if err != nil {
		if err == redis.Nil {
			c.AbortWithStatusJSON(404, ErrorResponse{"kvKey not found"})
			return
		}
		c.AbortWithStatusJSON(500, ErrorResponse{"error getting kvKey"})
		return
	}
	c.JSON(200, ResultResponse{kvValue})
}

func PutValue(c *gin.Context) {
	kvKey := formatKvKey(c.Param(kvKeyParam))
	kvValue := c.Query("value")
	if kvValue == "" {
		c.AbortWithError(400, fmt.Errorf("Missing value query param in PutValue"))
		return
	}
	err := rdb.Set(kvKey, kvValue, 0).Err()
	if err != nil {
		c.AbortWithStatusJSON(500, ErrorResponse{"error writing kvKey"})
		return
	}
	c.JSON(200, SuccessResponse{"success"})
}

func DeleteValue(c *gin.Context) {
	kvKey := formatKvKey(c.Param(kvKeyParam))
	err := rdb.Get(kvKey).Err()
	if err != nil {
		if err == redis.Nil {
			c.AbortWithStatusJSON(404, ErrorResponse{"kvKey not found"})
			return
		}
		c.AbortWithStatusJSON(500, ErrorResponse{"error getting kvKey"})
		return
	}
	if err = rdb.Del(kvKey).Err(); err != nil {
		c.AbortWithStatusJSON(500, ErrorResponse{"error deleting kvKey"})
		return
	}
	c.JSON(200, SuccessResponse{"success"})
}
