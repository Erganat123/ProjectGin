package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var user = map[string]string{
	"admin": "1234",
}

var jwtKey = []byte("secret")

// 🔐 LOGIN
func Login(c *gin.Context) {
	var creds map[string]string

	c.BindJSON(&creds)

	username := creds["username"]
	password := creds["password"]

	if user[username] != password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour).Unix(),
	})

	tokenString, _ := token.SignedString(jwtKey)

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// 🛡️ MIDDLEWARE (вот сюда вставляешь)
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenStr := c.GetHeader("Authorization")

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}