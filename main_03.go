package main

import (
	EmployeeController "backend/api/controller/employee"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//Employee API Method
	router.GET("/employees", EmployeeController.GET)
	router.POST("/employees", EmployeeController.POST)
	router.PUT("/employees", EmployeeController.PUT)
	router.DELETE("/employees", EmployeeController.DELETE)

	router.GET("/employees/:id", EmployeeController.GetEmployeeByID)       //GET By ID
	router.POST("/employees/:id", EmployeeController.PostEmployeeByID)     //GET By ID
	router.PUT("/employees/:id", EmployeeController.PUTEmployeeByID)       //GET By ID
	router.DELETE("/employees/:id", EmployeeController.DELETEEmployeeByID) //GET By ID

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
