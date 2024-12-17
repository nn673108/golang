package employee

import (
	"net/http"

	"backend/api/db"
	"github.com/gin-gonic/gin"
)

type Tbl_employee struct {
	Emp_id        int
	Emp_firstname string
	Emp_lastname  string
}

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
func GetEmployeeDB(c *gin.Context) {
	var employees []Tbl_employee
	db.Db.Find(&employees)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Employee Read Success", "employees": employees})
}

func PostEmployeeDB(c *gin.Context) {
	var employees []Tbl_employee
	db.Db.Find(&employees)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Employee Read Success", "employees": employees})
}

func PUTEmployeeDB(c *gin.Context) {
	var employees []Tbl_employee
	db.Db.Find(&employees)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Employee  Success", "employees": employees})
}
func DELETEEmployeeDB(c *gin.Context) {
	var employees []Tbl_employee
	db.Db.Find(&employees)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Employee Delete Success", "employees": employees})
}
