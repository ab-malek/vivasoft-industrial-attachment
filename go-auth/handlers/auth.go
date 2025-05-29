package handlers

import (
	"net/http"

	"go-auth/models"
	"go-auth/utils"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignUpHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req SignupRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
		}

		user := models.User{
			Name:     req.Name,
			Email:    req.Email,
			Password: req.Password, // simple for now; hash in production
		}

		if err := db.Create(&user).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Could not create user"})
		}

		return c.JSON(http.StatusOK, echo.Map{"message": "User created"})
	}
}

func LoginHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req LoginRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
		}

		var user models.User
		if err := db.Where("email = ? AND password = ?", req.Email, req.Password).First(&user).Error; err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid credentials"})
		}

		token, err := utils.GenerateToken(user.ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create token"})
		}

		return c.JSON(http.StatusOK, echo.Map{"token": token})
	}
}

func ProfileHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userID := uint(claims["user_id"].(float64))

		var u models.User
		if err := db.First(&u, userID).Error; err != nil {
			return c.JSON(http.StatusNotFound, echo.Map{"error": "User not found"})
		}

		return c.JSON(http.StatusOK, echo.Map{
			"id":    u.ID,
			"name":  u.Name,
			"email": u.Email,
		})
	}
}
