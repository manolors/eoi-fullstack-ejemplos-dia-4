package main

import (
	"initgin/middleware"
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

	const secretKey = "12345"

	// También se pueden poner por separado
	// r.Use(APIKey(secretKey))
	// r.Use(MiddlewareDummy)

	// middleware por función
	v1 := r.Group("/v1")
	{
		v1.POST("/nombre/:nombre", devuelveNombre)
		v1.POST("/apellidos/:apellidos", devuelveApellidos, middleware.Dummy)
	}

	// middleware para un grupo de rutas
	v2 := r.Group("/v2", middleware.APIKey(secretKey), middleware.Dummy)
	{
		v2.POST("/nombre/:nombre", devuelveNombre)
		v2.POST("/apellidos/:apellidos", devuelveApellidosV2)
	}
	r.Run(":8888")
}
