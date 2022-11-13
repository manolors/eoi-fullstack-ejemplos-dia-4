package main

import (
	"github.com/gin-gonic/gin"
)

func devuelveNombre(c *gin.Context) {
	user := c.Params.ByName("name")
	c.String(200, "%s", "hola "+user)
}

func main() {
	r := gin.Default()

	// Get user value
	r.GET("/nombre/:name", devuelveNombre)

	r.Run(":8888")

}
