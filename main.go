package main

import (
	"go-demo/mac"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/mac", func(c *gin.Context) {
		mac, err := mac.Fetch()
		if err != nil {
			c.JSON(200, gin.H{"err": err})
			return
		}
		c.JSON(200, gin.H{"data": mac})
	})

	router.Run("127.0.0.1:8080")
}
