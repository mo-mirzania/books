package application

import (
	"log"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

// StartApp func
func StartApp() {
	Mappings()
	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
