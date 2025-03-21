package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/LeoneIAguilera/web-simple-two/initializers"
	"github.com/LeoneIAguilera/web-simple-two/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	// Get the cookie off req
	tokenString, err := c.Cookie("Authorization")
	
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Decode/validate it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
	
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// Check exp
		expTime := claims["exp"].(float64)
		nowTime := float64(time.Now().Unix())

		if nowTime > expTime {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		
		// Find the user token sub
		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Attach to req
		c.Set("user", user)		

		// Continue
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}