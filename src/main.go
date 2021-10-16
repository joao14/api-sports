package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello Worold",
	})
}

func main() {
	r := gin.Default()

	r.GET("/", HomePage)

	r.POST("/player/:name/*action", func(c *gin.Context) {
		b := c.FullPath() == "/player/:name/*action"
		c.String(http.StatusOK, "%t", b)
	})

	r.Run(":8082")
}
