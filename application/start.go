package application

import (
	"log"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

// StartApp func
func StartApp() {
	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()
	Mappings()
	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
