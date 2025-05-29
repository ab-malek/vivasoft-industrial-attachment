package initializers

import "go-jwt-auth/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}