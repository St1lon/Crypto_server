package handlers

import (
	//"cryptoserver/domain"
	"cryptoserver/errors"
	"cryptoserver/internal/repository"
	"cryptoserver/internal/service"
	"encoding/json"
	"log"
	"net/http"
	//"github.com/golang-jwt/jwt/v5"
	//"cryptoserver/middleware"
)

func HandlerAuth(userRepo repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			customErr := errors.NewErrWrongMethod("wrong method: "+r.Method, http.StatusMethodNotAllowed, "login user")
			WriteJsonError(w, customErr)
			log.Println(customErr)
			return
		}
		if ct := r.Header.Get("Content-Type"); ct != "application/json" {
			customErr := errors.NewErrWrongCT("Content-Type must be application/json", http.StatusUnsupportedMediaType, "login user")
			WriteJsonError(w, customErr)
			log.Println("unsupported content type:", customErr)
			return
		}
		var user_request RegisterRequest
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&user_request)
		if err != nil {
			customErr := errors.NewErrInvalidJSON(err.Error(), http.StatusBadRequest, "login user")
			WriteJsonError(w, customErr)
			log.Printf("JSON decode error: %v", err)
			return
		}
		if user_request.Username == "" {
			customErr := errors.NewErrUserNameRequired("username is required field", http.StatusBadRequest, "login user")
			WriteJsonError(w, customErr)
			log.Println(customErr)
			return
		}
		if user_request.Password == "" {
			customErr := errors.NewErrPasswordRequired("password is required field", http.StatusBadRequest, "login user")
			WriteJsonError(w, customErr)
			log.Println(customErr)
			return
		}
		authService := service.NewAuthService(userRepo)
		token, customErr := authService.Login(user_request.Username, user_request.Password)
		if customErr != nil {
			WriteJsonError(w, customErr)
			log.Printf("%s", customErr.Error())
			return
		}
		response := map[string]string{"token": token}
		WriteJsonResponse(w, response, http.StatusOK)
		log.Printf("user '%s' logged in successfully", user_request.Username)
	}
}
