package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// empployee API Metod
func main() {
	r := gin.Default()
	r.GET("/employees", getemployee)
	r.POST("/employees", postemployee)
	r.PUT("/employees", putemployee)
	r.DELETE("/employees", deleteemployee)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getemployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Employee GET Method",
	})
}

func postemployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Employee GET Method",
	})
}

func putemployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Employee GET Method",
	})
}

func deleteemployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Employee GET Method",
	})
}
