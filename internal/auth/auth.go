package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// NewJWT takes
func NewJWT(issuer string) *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Local().Add(time.Minute * 30).Unix(),
		IssuedAt:  time.Now().Local().Unix(),
		Issuer:    issuer,
	})
}
