package handler

import (
	//"cryptoserver/domain"
	"cryptoserver/errors"
	"cryptoserver/repository"
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
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
		user, err := userRepo.GetByUsername(user_request.Username)
		if err != nil {
			customErr := errors.NewErrUserNotFound("user not found", http.StatusNotFound, "login user")
			WriteJsonError(w, customErr)
			log.Printf("%s: %v", customErr.Error(), err)
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(user_request.Password))
		if err != nil {
			customErr := errors.NewErrInvalidCredentials("invalid credentials", http.StatusUnauthorized, "login user")
			WriteJsonError(w, customErr)
			log.Printf("%s: %v", customErr.Error(), err)
			return
		}
		token, err := GenerateToken(user)
		if err != nil {
			customErr := errors.NewErrGenerateToken("failed to generate token", http.StatusInternalServerError, "login user")
			WriteJsonError(w, customErr)
			log.Printf("%s: %v", customErr.Error(), err)
			return
		}
		response := map[string]string{"token": token}
		WriteJsonResponse(w, response, http.StatusOK)
		log.Printf("user '%s' logged in successfully", user_request.Username)
	}
}
