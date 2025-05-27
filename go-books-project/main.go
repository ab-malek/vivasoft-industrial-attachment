package main

import (
    "go-books-project/db"
    "go-books-project/handlers"

    "github.com/labstack/echo/v4"
)

func main() {
    db.Init()

    e := echo.New()

    e.POST("/books", handlers.CreateBook)
    e.GET("/books", handlers.GetAllBooks)
    e.PUT("/books/:id", handlers.UpdateBook)
    e.DELETE("/books/:id", handlers.DeleteBook)

    e.Logger.Fatal(e.Start(":8080"))
}
