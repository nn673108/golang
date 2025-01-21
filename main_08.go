package main

import (
	AdminController "backend/api/controller/admin"
	AuthController "backend/api/controller/auth"
	EmployeeController "backend/api/controller/employee"
	"backend/api/db"
	"backend/api/middleware" //เรียกใช้ไฟล์ที่อยู่ในห้อง middleware
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

	authorized := router.Group("/api", middleware.JwtAuthen()) //ทำการจัดกลุ่ม path ที่ต้องการล๊อค api

	//Employee API Method
	router.GET("/employee", EmployeeController.GetEmployee)         //GET
	router.GET("/employee/:id", EmployeeController.GetEmployeeByID) //GET BY ID
	authorized.GET("/employeedb", EmployeeController.GetEmployeeDB) //ล๊อค api โดยต้องแนบ token ก่อนถึงใช้งานได้
	router.GET("/admin", AdminController.GetAdmin)                  //POST TO DB

	router.POST("/employee", EmployeeController.PostEmployee)     //POST TO DB
	router.POST("/employeedb", EmployeeController.PostEmployeeDB) //POST TO DB

	router.POST("/register", AdminController.PostAdmin) //POST TO DB

	router.PUT("/employee", EmployeeController.PutEmployee)     //PUT
	router.PUT("/employeedb", EmployeeController.PutEmployeeDB) //PUT DB

	router.DELETE("/employee", EmployeeController.DeleteEmployee)         //DELETE
	router.DELETE("/employeedb/:id", EmployeeController.DeleteEmployeeDB) //DELETE DB

	//Customer API Method
	router.POST("/login", AuthController.Login) //POST LOGIN
	router.Run()                                // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
