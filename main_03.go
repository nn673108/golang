package main

import (
	EmployeeController "backend/api/controller/employee"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//Employee API Method
	router.GET("/employee", EmployeeController.GET)
	router.POST("/employee", EmployeeController.POST)
	router.PUT("/employee", EmployeeController.PUT)
	router.DELETE("/employee", EmployeeController.DELETE)

	router.GET("/employee/:id", EmployeeController.GetEmployeeByID)       //GET By ID
	router.POST("/employee/:id", EmployeeController.PostEmployeeByID)     //GET By ID
	router.PUT("/employee/:id", EmployeeController.PUTEmployeeByID)       //GET By ID
	router.DELETE("/employee/:id", EmployeeController.DELETEEmployeeByID) //GET By ID

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
