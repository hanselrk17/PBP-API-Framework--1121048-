package main

import (
	"fmt"

	"log"
	"modul-eksplorasi/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	router := gin.Default()
	router.GET("/users_gorm", controllers.GetAllUsersGorm)
	router.GET("/users", controllers.GetAllUsers)

	router.POST("/user_gorm", controllers.InsertUserGORM)
	router.POST("/user", controllers.InsertUser)

	router.PUT("/user/:userId", controllers.UpdateUser)

	router.DELETE("/user/:userId", controllers.DeleteUser)

	http.Handle("/", router)
	fmt.Println("Connected to port 6060")
	log.Println("Connected to port 6060")
	router.Run("localhost:6060")
}
