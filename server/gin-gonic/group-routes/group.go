package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func devuelveNombre(c *gin.Context) {
	nombre := c.Params.ByName("nombre")
	c.JSON(http.StatusOK, gin.H{"nombre": nombre, "version": "v1"})
}
func devuelveApellidos(c *gin.Context) {
	apellidos := c.Params.ByName("apellidos")
	c.JSON(http.StatusOK, gin.H{"apellidos": apellidos, "version": "v1"})
}

func devuelveApellidosV2(c *gin.Context) {
	apellidos := c.Params.ByName("apellidos")
	c.JSON(http.StatusOK, gin.H{"nombre": apellidos, "version": "v2"})
}

func main() {
	r := gin.Default()

	v1 := r.Group("/nombre")
	{
		v1.DELETE("nombre/:nombre", devuelveNombre)
		v1.GET("/:nombre", devuelveNombre)
	}

	v2 := r.Group("/v2")
	{
		v2.POST("/nombre/:nombre", devuelveNombre)
		v2.POST("/apellidos/:apellidos", devuelveApellidosV2)
	}
	r.Run(":8888")
}
