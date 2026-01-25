package middleware

import (
	//"crypto/aes"
	"fmt"
	"cryptoserver/errors"
	"log"
	"net/http"
	"cryptoserver/handler"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			customErr := errors.NewErrTokenMissed("Authorization token is missing", http.StatusUnauthorized, "auth middleware")
			handler.WriteJsonError(w, customErr)
			log.Println(customErr)
			return
		}
		token, err := handler.ParseToken(tokenString)
		if err != nil || !token.Valid {
			customErr := errors.NewErrInvalidToken("Invalid authorization token", http.StatusUnauthorized, "auth middleware")
			handler.WriteJsonError(w, customErr)
			log.Println(fmt.Errorf("%s : %w", &customErr, err))
			return
		}
		handler.WriteJsonResponse(w, map[string]string{"message": tokenString}, http.StatusOK)
		next.ServeHTTP(w, r)
	})
}