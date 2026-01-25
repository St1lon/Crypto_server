package handler

import (
	"cryptoserver/domain"
	"cryptoserver/errors"
	"cryptoserver/repository"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	//"github.com/golang-jwt/jwt/v5"
	"cryptoserver/middleware"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"` // Приходит password, не hash!
}

func HandlerRegister(userRepo repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			customErr := errors.NewErrWrongMethod("wrong method: "+r.Method, http.StatusMethodNotAllowed, "register user")
			WriteJsonError(w, customErr)
			log.Println(customErr)
			return
		}

		if ct := r.Header.Get("Content-Type"); ct != "application/json" {
			customErr := errors.NewErrWrongCT("Content-Type must be application/json", http.StatusUnsupportedMediaType, "register user")
			WriteJsonError(w, customErr)
			log.Println("unsupported content type:", customErr)
			return
		}
		var user_request RegisterRequest
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&user_request)
		if err != nil {
			customErr := errors.NewErrInvalidJSON(err.Error(), http.StatusBadRequest, "register user")
			WriteJsonError(w, customErr)
			log.Printf("JSON decode error: %v", err)
			return
		}
		if user_request.Username == "" {
			customErr := errors.NewErrUserNameRequired("username is required field", http.StatusBadRequest, "register user")
			WriteJsonError(w, customErr)
			log.Println(customErr)
			return
		}
		if user_request.Password == "" {
			customErr := errors.NewErrPasswordRequired("password is required field", http.StatusBadRequest, "register user")
			WriteJsonError(w, customErr)
			log.Println(customErr)
			return
		}

		var user domain.User
		_, err = userRepo.GetByUsername(user_request.Username)
		if err == nil {
			customErr := errors.NewErrUserAlreadyExists("user with this username already exists", http.StatusConflict, "register user")
			WriteJsonError(w, customErr)
			log.Printf("user registration conflict: username '%s' already exists", user_request.Username)
			return
		}
		user.Username = user_request.Username
		hash, err := bcrypt.GenerateFromPassword([]byte(user_request.Password), bcrypt.DefaultCost)
		if err != nil {
			customErr := errors.NewErrHashingPassword("fail to hash password", http.StatusInternalServerError, "register user")
			WriteJsonError(w, customErr)
			log.Printf("bcrypt hashing error: %v", err)
			return
		}
		user.PasswordHash = string(hash)
		err = userRepo.Create(&user)
		if err != nil {
			customErr := errors.NewErrCreateUser("fail to create user", http.StatusInternalServerError, "register user")
			WriteJsonError(w, customErr)
			log.Printf("user creation error: %v", err)
			return
		}

		token, err := middleware.GenerateToken(&user)
		if err != nil {
			customErr := errors.NewErrGenerateToken("fail to generate JWT token", http.StatusInternalServerError, "register user")
			WriteJsonError(w, customErr)
			log.Printf("token generation error: %v", err)
			return
		}
		WriteJsonResponse(w, map[string]string{
			"token": token,
		}, http.StatusCreated)
		log.Println("user registered:", user.Username)
	}
}

func WriteJsonError(w http.ResponseWriter, err errors.CustomError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.GetCode())
	json.NewEncoder(w).Encode(map[string]string{
		"error": err.GetMsg() + " Op:" + err.GetOp() + " Code:" + fmt.Sprintf("%d", err.GetCode()),
	})
}
func WriteJsonResponse(w http.ResponseWriter, message map[string]string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(message)
}
