package handler

import (
	"github.com/kayquesza/api-go-learning/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
	// db *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}
