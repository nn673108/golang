package employee

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET
func GET(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Employee GET Method!",
	})
}
func POST(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Employee Post Method!",
	})
}
func PUT(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Employee Put Method!",
	})
}
func DELETE(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Employee Delete Method!",
	})
}
