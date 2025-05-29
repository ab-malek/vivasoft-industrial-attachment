package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo-jwt/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"go-auth/models"
	"go-auth/handlers"
	"go-auth/utils"
)

func main() {
	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database")
	}

	db.AutoMigrate(&models.User{})

	e := echo.New()

	// Public routes
	e.POST("/signup", handlers.SignUpHandler(db))
	e.POST("/login", handlers.LoginHandler(db))

	// JWT Middleware group
	profileGroup := e.Group("/profile")
	profileGroup.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: utils.JwtSecret,
	}))

	// Protected route
	profileGroup.GET("", handlers.ProfileHandler(db))

	e.Logger.Fatal(e.Start(":8080"))
}
