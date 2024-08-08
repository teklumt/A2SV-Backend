package infrastracture

import (
	"clean_architecture/domain"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(user domain.User) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user.ID
	claims["role"] = user.Role
	claims["username"] = user.Username


	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}


