package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func APIKey(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("Api-Key")
		if secretKey != key {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"status":  http.StatusForbidden,
				"message": "Permission denied",
			})
			return
		}
		c.Next()
	}
}
