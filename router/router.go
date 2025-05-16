package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {

	// Inicializa o Router utilizando as configurações Default do Gin
	router := gin.Default()

	// Define uma rota GET para o endpoint "/ping"
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8080") // listen and serve on 0.0.0.0:8080

}
