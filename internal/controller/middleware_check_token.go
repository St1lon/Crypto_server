package controller

import (
	"cryptoserver/internal/domain"
	"cryptoserver/internal/service"
	"log"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			customErr := domain.NewErrTokenMissed("Authorization token is missing", http.StatusUnauthorized, "auth middleware")
			WriteJsonError(w, customErr)
			log.Println(customErr)
			return
		}
		token, err := service.ParseToken(tokenString)
		if err != nil || !token.Valid {
			customErr := domain.NewErrInvalidToken("Invalid authorization token", http.StatusUnauthorized, "auth middleware")
			WriteJsonError(w, customErr)
			if err != nil {
				log.Printf("%s: %v", customErr.Error(), err)
			} else {
				log.Println(customErr.Error())
			}
			return
		}
		next.ServeHTTP(w, r)
	})
}
