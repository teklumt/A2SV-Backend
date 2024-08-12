package infrastracture

import (
	"clean_architecture_Testing/config"
	"clean_architecture_Testing/domain"

	"github.com/golang-jwt/jwt"
)


func GenerateToken(user domain.User) (string, error) {
	jwtSecret := []byte(config.EnvConfigs.JwtSecret)


	claims := domain.JwtCustomClaims{
		Authorized: true,
		UserID:     user.ID.Hex(),
		Role:       user.Role,
		Username:   user.Username,
	}


	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}


