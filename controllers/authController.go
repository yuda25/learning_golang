package controllers

import (
	"learning_golang/initializers"
	"learning_golang/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	// get data from body
	var body struct {
		Name string `json:"name"`
		Email string `json:"email" gorm:"unique"`
		Password string `json:"password"`
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "failed to read body",
		})
		return
	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "failed to hash password",
		})
		return
	}

	// create user
	user := models.User{
		Name: body.Name,
		Email: body.Email,
		Password: string(hash),
	}
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "failed to create user",
		})
		return
	}
	// response
	c.JSON(http.StatusOK, gin.H{
		"message": "signup successfully",
	})
}

func SignIn(c *gin.Context)  {
	// get fata
	var body struct {
		Email string `json:"email" gorm:"unique"`
		Password string `json:"password"`
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "failed to read body",
		})
		return
	}

	// find for DB
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "Invalid email or password",
		})
		return
	}

	// validasi password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "Invalid email or password",
		})
		return
	}

	// generate JWT token
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"name": user.Name,
		"email": user.Email,
		"exp":time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "failed to create token",
		})
		return
	}

	// return
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		// "token": tokenString,
		"message": "login successfully",
	})
}

func Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":"you are logged in",
	})
}