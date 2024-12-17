package employee

import (
	"net/http"

	"backend/api/db"

	"github.com/gin-gonic/gin"
)

type Tbl_employee struct {
	Emp_id         int
	Emp_firstname  string
	Emp_lastname   string
	Emp_salary     float64
	Emp_department string
}

// Biding from Employee JSON
type EmployeeBody struct {
	Emp_id         int     `json:"emp_id" binding:"required"`
	Emp_firstname  string  `json:"emp_firstname" binding:"required"`
	Emp_lastname   string  `json:"emp_lastname" binding:"required"`
	Emp_department string  `json:"emp_department" binding:"required"`
	Emp_salary     float64 `json:"emp_salary" binding:"required"`
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

// get DB
func GetEmployeeDB(c *gin.Context) {
	var employees []Tbl_employee
	db.Db.Find(&employees)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Employee Read Success", "employees": employees})
}

// POST Employee to Database
func PostEmployeeDB(c *gin.Context) {
	var json EmployeeBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tbl_employee := Tbl_employee{Emp_id: json.Emp_id, Emp_firstname: json.Emp_firstname, Emp_lastname: json.Emp_lastname, Emp_department: json.Emp_department, Emp_salary: json.Emp_salary}
	db.Db.Create(&tbl_employee)
	if tbl_employee.Emp_firstname != "" {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Employee Created", "tbl_fund": tbl_employee})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "User Failed", "tbl_fund": tbl_employee})
	}

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
