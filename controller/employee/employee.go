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

// GET By ID
func GetEmployeeByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "GET GET" + id,
	})
}

func PostEmployeeByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "POST POST" + id,
	})
}

//GET By ID

func PUTEmployeeByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "PUT PUT" + id,
	})
}

func DELETEEmployeeByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "get out" + id,
	})
}
