package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Dummy(c *gin.Context) {
	log.Println("Pasamos por el middleware!")
	c.Next()
}
