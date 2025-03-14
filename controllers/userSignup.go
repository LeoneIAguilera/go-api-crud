package controllers

import (
	"net/http"

	"github.com/LeoneIAguilera/web-simple-two/initializers"
	"github.com/LeoneIAguilera/web-simple-two/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var body struct {
		Email string
		Password string
	}
	
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to read body",
		})
		
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Fail to hash password",
		})
		
		return
	}

	newUser := models.User{
		Email: body.Email,
		Password: string(hash),
	}

	result := initializers.DB.Create(&newUser)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Fail to create new user",
		})
		
		return
	}
	
	c.JSON(http.StatusOK, gin.H{})

}