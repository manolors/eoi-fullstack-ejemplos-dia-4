package main

import "github.com/gin-gonic/gin"

func handler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
func main() {
	r := gin.Default()
	r.GET("/ping", handler)
	r.Run(":8989")
}
