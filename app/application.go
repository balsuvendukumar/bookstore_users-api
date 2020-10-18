package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	fmt.Println("In App1")
	MapUrls()
	router.Run(":8080")
}
