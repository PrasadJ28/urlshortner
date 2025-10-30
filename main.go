package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Printf("Hello Go URL Shortner!")

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Go URL Shortener!",
		})
	})

	err := r.Run(":9808")
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
