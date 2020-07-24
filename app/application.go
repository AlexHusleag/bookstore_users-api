package app

import (
	"github.com/gin-gonic/gin"
)

// HTTP only here and in controller

var(
	router = gin.Default()
)

func StartApplication(){
	mapURLs()
	router.Run(":8080")
}
