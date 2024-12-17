package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// empployee API Metod
func main() {
	r := gin.Default()
	r.GET("/employees", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Informaion Technology",
		})
	})
	//Default API Method

	r.POST("/employees", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "I top1 err",
		})
	})
	r.PUT("/employees", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "I top err",
		})
	})
	r.DELETE("/employees", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "I top err2",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
