package infrastracture

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			c.JSON(401, gin.H{"error": "Invalid authorization header"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})

		if err != nil {
			fmt.Printf("Error parsing token: %v\n", err)
			c.JSON(401, gin.H{"error": "Invalid JWT"})
			c.Abort()
			return
		}

		if !token.Valid {
			fmt.Println("Token is not valid")
			c.JSON(401, gin.H{"error": "Invalid JWT"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(401, gin.H{"error": "Invalid JWT"})
			c.Abort()
			return
		}

		c.Set("user", claims)
		// c.Set("Roles", claims["role"])
		c.Next()
	}
}






func RoleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userClaims, exists := c.Get("user")
		// fmt.Println(userClaims)


		if !exists {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		user, ok := userClaims.(jwt.MapClaims)
		if !ok {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user claims"})
			c.Abort()
			return
		}

		role, ok := user["role"].(string)
		if !ok || role == "" {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Role not found in claims"})
			c.Abort()
			return
		}

		if role != "admin" {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
			c.Abort()
			return
		}

		c.Next()
	}
}

