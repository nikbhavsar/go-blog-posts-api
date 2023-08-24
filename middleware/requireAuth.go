package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nikbhavsar/go-blog-posts-api/initalizers"
	"github.com/nikbhavsar/go-blog-posts-api/models"
)

func RequireAuth(c *gin.Context) {
	//Get the coockie
	tokenString, err := c.Cookie("auth")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	//Decode and validate the token

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//Check the expiry
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		//Find the user with token
		var user models.User
		initalizers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		//Attach to request
		c.Set("user", user)
		//contine
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
