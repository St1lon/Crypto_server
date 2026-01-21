package middleware

import (
	"cryptoserver/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
)
var secret_key = []byte("clmcmekkenckmekme")

func GenerateToken(user *domain.User) (string, error){
	claims := jwt.MapClaims{
		"username" : user.Username,
		"exp" : time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret_key)
}
