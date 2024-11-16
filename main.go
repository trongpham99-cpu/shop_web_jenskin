package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/api/v1/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "get all products",
		})
	})
	r.GET("/api/v1/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": "get product by id",
			"id":      id,
		})
	})
	r.Run(":3000")
}
