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
	router.GET("/employees", EmployeeController.GET)
	router.POST("/employees", EmployeeController.POST)
	router.PUT("/employees", EmployeeController.PUT)
	router.DELETE("/employees", EmployeeController.DELETE)

	router.GET("/employees/:id", EmployeeController.GetEmployeeByID)       //GET By ID
	router.POST("/employees/:id", EmployeeController.PostEmployeeByID)     //GET By ID
	router.PUT("/employees/:id", EmployeeController.PUTEmployeeByID)       //GET By ID
	router.DELETE("/employees/:id", EmployeeController.DELETEEmployeeByID) //GET By ID

	router.GET("/employeesdb", EmployeeController.GetEmployeeDB)       //GET DB
	router.POST("/employeesdb", EmployeeController.PostEmployeeDB)     //GET DB
	router.PUT("/employeesdb", EmployeeController.PUTEmployeeDB)       //GET DB
	router.DELETE("/employeesdb", EmployeeController.DELETEEmployeeDB) //GET DB

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
