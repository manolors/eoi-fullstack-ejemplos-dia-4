package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func procesarPersonaMust(c *gin.Context) {
	type persona struct {
		Nombre    string `form:"name" json:"name" binding:"required"`
		Apellidos string `form:"apellidos" json:"surname" binding:"required"`
	}

	var p persona
	err := c.MustBindWith(&p, binding.JSON)

	if err != nil {
		return
	}

	c.JSON(200, gin.H{"mensaje": "Todo ok"})

}
func procesarPersonaShould(c *gin.Context) {
	type persona struct {
		Nombre    string `json:"name" binding:"required"`
		Apellidos string `json:"surname" binding:"required"`
	}

	var p persona
	c.ShouldBindWith(&p, binding.JSON)

	c.JSON(200, gin.H{"mensaje": "Todo ok"})
}

func main() {
	r := gin.Default()

	// Get user value
	r.POST("/must/", procesarPersonaMust)
	r.POST("/should/", procesarPersonaShould)

	r.Run(":8888")

}
