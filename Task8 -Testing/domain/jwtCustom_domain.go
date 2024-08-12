package domain

import (
	"github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaims struct {
	Authorized bool   `json:"authorized"`
	UserID     string `json:"user_id"`
	Role       string `json:"role"`
	Username   string `json:"username"`
	
	jwt.StandardClaims
}


