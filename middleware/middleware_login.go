package middleware

import (
	//"crypto/aes"

	"cryptoserver/errors"
	"cryptoserver/handler"
	"log"
	"net/http"
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
			if err != nil {
				log.Printf("%s: %v", customErr.Error(), err)
			} else {
				log.Println(customErr.Error())
			}
			return
		}
		//handler.WriteJsonResponse(w, map[string]string{"message": tokenString}, http.StatusOK)
		next.ServeHTTP(w, r)
	})
}
