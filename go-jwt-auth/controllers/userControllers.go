package controllers

import (
	"go-jwt-auth/initializers"
	"go-jwt-auth/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context){

	// get data from request
	var body struct{
		Email string
		Password string
	}

	if c.Bind(&body) != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	// hash the passwrod

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	// create the user
	user := models.User{Email: body.Email, Passwrod: string(hash)}

	result := initializers.DB.Create(&user)

	if result.Error != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}


func Login(c *gin.Context){
	// get the data
	var body struct{
		Email string
		Password string
	}

	if c.Bind(&body) != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// check is data exist
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if(user.ID == 0){
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}
	// compare is password

	err := bcrypt.CompareHashAndPassword([]byte(user.Passwrod), []byte(body.Password))

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"expire": time.Now().Add(time.Hour*24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Can't create token",
		})
		return
	}

	c.SetSameSite(http.SameSiteDefaultMode)
	c.SetCookie("Authorization", tokenString, 36000, "", "", false,true)

	// c.JSON(http.StatusOK, gin.H{
	// 	"Token" : tokenString,
	// })

	// send the token
}

func Validate(c *gin.Context){
	user,_ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"User" : user,
	})
}