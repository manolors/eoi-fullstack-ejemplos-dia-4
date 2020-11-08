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

	// middleware global
	const secretKey = "12345"

	// También se pueden poner por separado
	// r.Use(APIKey(secretKey))
	// r.Use(MiddlewareDummy)

	v1 := r.Group("/v1")
	{
		v1.POST("/nombre/:nombre", devuelveNombre)
		v1.POST("/apellidos/:apellidos", devuelveApellidos, MiddlewareDummy)
	}

	v2 := r.Group("/v2", APIKey(secretKey), MiddlewareDummy) // middleware para el grupo
	{
		v2.POST("/nombre/:nombre", devuelveNombre)
		v2.POST("/apellidos/:apellidos", devuelveApellidosV2)
	}
	r.Run(":8888")
}
