package main

import (
	"github.com/kayquesza/api-go-learning/config"
	_ "github.com/kayquesza/api-go-learning/docs"
	"github.com/kayquesza/api-go-learning/router"
)

// @title           API Go Learning
// @version         1.0
// @description     API de estudo para Swagger em Go
// @host            localhost:8080
// @BasePath        /api/v1


var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
