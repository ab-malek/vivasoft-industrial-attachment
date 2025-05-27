package handlers

import (
    "net/http"
    "strconv"

    "github.com/labstack/echo/v4"
    "go-books-project/db"
    "go-books-project/models"
)

func CreateBook(c echo.Context) error {
    var book models.Book
    if err := c.Bind(&book); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }

    if err := db.DB.Create(&book).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusCreated, book)
}

func GetAllBooks(c echo.Context) error {
    var books []models.Book
    if err := db.DB.Find(&books).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusOK, books)
}

func UpdateBook(c echo.Context) error {
    id := c.Param("id")
    var book models.Book

    if err := db.DB.First(&book, id).Error; err != nil {
        return c.JSON(http.StatusNotFound, "Book not found")
    }

    var updated models.Book
    if err := c.Bind(&updated); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }

    book.Name = updated.Name
    book.Author = updated.Author
    book.Description = updated.Description

    if err := db.DB.Save(&book).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusOK, book)
}

func DeleteBook(c echo.Context) error {
    id := c.Param("id")
    bookID, _ := strconv.Atoi(id)

    if err := db.DB.Delete(&models.Book{}, bookID).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusOK, "Book deleted successfully")
}
