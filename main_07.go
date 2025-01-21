package main

import (
	EmployeeController "backend/api/controller/employee"
	"backend/api/db"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//Get .env
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	//get InitDB fuction
	db.InitDB()

	router := gin.Default()
	//Employee API Method
	router.GET("/employees", EmployeeController.GetEmployee)         //GET
	router.GET("/employees/:id", EmployeeController.GetEmployeeByID) //GET BY ID
	router.GET("/employeesdb", EmployeeController.GetEmployeeDB)     //GET FROM DB

	router.POST("/employee", EmployeeController.PostEmployee)     //POST
	router.POST("/employeedb", EmployeeController.PostEmployeeDB) //POST TO DB

	router.PUT("/employee", EmployeeController.PutEmployee)     //PUT
	router.PUT("/employeedb", EmployeeController.PutEmployeeDB) //PUT TO DB

	router.DELETE("/employee", EmployeeController.DeleteEmployee)         //DELETE
	router.DELETE("/employeedb/:id", EmployeeController.DeleteEmployeeDB) //DELETE DB

	//Customer API Method

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
