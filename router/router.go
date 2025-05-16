package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {

	// Initialize Router
	router := gin.Default()

	// Initialize Routes
	initializeRoutes(router)

	// Run the server on port 8080
	router.Run(":8080") // Listen and serve on 0.0.0.0:8080

}
