package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func procesarPersonaShould(c *gin.Context) {
	type persona struct {
		Nombre    string `json:"name"`
		Apellidos string `json:"surname"`
	}

	var p persona
	c.ShouldBindWith(&p, binding.JSON)
	log.Printf("%#v", p)
	c.JSON(200, gin.H{"mensaje": "Todo ok"})
}

func main() {
	r := gin.Default()

	// Get user value
	r.GET("/nombre/:name", procesarPersonaShould)

	r.Run(":8888")

}
