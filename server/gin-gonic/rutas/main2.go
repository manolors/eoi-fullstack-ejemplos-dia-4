package main

import (
	"github.com/gin-gonic/gin"
)

func devuelveNombre(c *gin.Context) {
	//	user := c.Params.ByName("name")
	// c.JSON(http.StatusOK, gin.H{"nombre": user})
	//	c.HTML(200, "nombre", "<html><h1>hola</h1></html>")
	c.String(200, "%s", "hola")
}

func main() {
	r := gin.Default()

	// Get user value
	r.GET("/nombre/:name", devuelveNombre)

	r.Run(":8888")

}
