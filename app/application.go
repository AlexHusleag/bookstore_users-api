package app

import (
	"github.com/AlexHusleag/bookstore_users-api/logger"
	"github.com/gin-gonic/gin"
)

// HTTP only here and in controller

var (
	router = gin.Default()
)

func StartApplication() {
	mapURLs()
	logger.Info("About to start the application...")
	router.Run(":8080")
}
