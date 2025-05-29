package main

import (
	"fmt"
	"go-jwt-auth/controllers"
	"go-jwt-auth/initializers"

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


	router.Run() // listen and serve on 0.0.0.0:8080
}