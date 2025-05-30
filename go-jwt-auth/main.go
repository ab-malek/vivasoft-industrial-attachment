package main

import (
	"fmt"
	"go-jwt-auth/controllers"
	"go-jwt-auth/initializers"
	"go-jwt-auth/middleware"

	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main(){
	fmt.Println("heloo sadamleasasdfk ")
	router := gin.Default()
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.GET("/login", middleware.RequireAuth, controllers.Validate)

	router.Run() // listen and serve on 0.0.0.0:8080
}