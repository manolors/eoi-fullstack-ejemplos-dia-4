package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func MiddlewareDummy(c *gin.Context) {
	log.Println("Pasamos por el middleware!")
	c.Next()
}
