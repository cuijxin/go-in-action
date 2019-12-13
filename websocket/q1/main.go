package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	bindAddress := "localhost:2303"
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, []string{"123", "321"})
	})
	r.Run(bindAddress)
}
