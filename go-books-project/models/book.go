package models

type Book struct {
    ID          uint   `gorm:"primaryKey" json:"id"`
    Name        string `json:"name"`
    Author      string `json:"author"`
    Description string `json:"description"`
}
