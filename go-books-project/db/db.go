package db

import (
	"go-books-project/models"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

var DB *gorm.DB

func Init() {
    var err error
    DB, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database: ")
    }

    // Auto migrate Book struct to create/update the table
    err = DB.AutoMigrate(&models.Book{})
    if err != nil {
        panic("failed to migrate database:")
    }
}

