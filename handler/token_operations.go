package handler

import (
	"cryptoserver/domain"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"fmt"
)
var secret_key = []byte("clmcmekkenckmekme")

type CustomClaims struct {
    Username string `json:"username"`
    jwt.RegisteredClaims // встраивание стандартных claims
}

func GenerateToken(user *domain.User) (string, error){
	claims := CustomClaims{
    Username: user.Username,
    RegisteredClaims: jwt.RegisteredClaims{
        ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
        IssuedAt:  jwt.NewNumericDate(time.Now()),
    },
}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret_key)
}

func ParseToken(tokenString string) (*jwt.Token, error) {
    return jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return secret_key, nil
    })
}