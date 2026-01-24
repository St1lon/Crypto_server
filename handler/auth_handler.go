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
			WriteJsonError(w, errors.NewErrWrongMethod(r.Method, http.StatusMethodNotAllowed, "register user"))
			log.Println("wrong method:", r.Method)
			return
		}

		if ct := r.Header.Get("Content-Type"); ct != "application/json" {
			WriteJsonError(w, errors.NewErrWrongCT(ct, http.StatusUnsupportedMediaType, "register user"))
			log.Println("unsupported content type:", ct)
			return
		}
		var user_request RegisterRequest
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&user_request)
		if err != nil {
			customErr := errors.NewErrInvalidJSON(err.Error(), http.StatusBadRequest, "register user")
			WriteJsonError(w, customErr)
			wrappedErr := fmt.Errorf("%s: %w", customErr, err)
			log.Println(wrappedErr)
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
			wrappedErr := fmt.Errorf("%s: %w", customErr, err)
			log.Println(wrappedErr)
			return
		}
		user.Username = user_request.Username
		hash, err := bcrypt.GenerateFromPassword([]byte(user_request.Password), bcrypt.DefaultCost)
		if err != nil {
			customErr := errors.NewErrHashingPassword("fail to hash password", http.StatusInternalServerError, "register user")
			WriteJsonError(w, customErr)
			wrappedErr := fmt.Errorf("%s: %w", customErr, err)
			log.Println(wrappedErr)
			return
		}
		user.PasswordHash = string(hash)
		err = userRepo.Create(&user)
		if err != nil {
			customErr := errors.NewErrCreateUser("fail to create user", http.StatusInternalServerError, "register user")
			WriteJsonError(w, customErr)
			wrappedErr := fmt.Errorf("%s: %w", customErr, err)
			log.Println(wrappedErr)
			return
		}

		token, err := middleware.GenerateToken(&user)
		if err != nil {
			customErr := errors.NewErrGenerateToken("fail to generate JWT token", http.StatusInternalServerError, "register user")
			WriteJsonError(w, customErr)
			wrappedErr := fmt.Errorf("%s: %w", customErr, err)
			log.Println(wrappedErr)
			return
		}
		WriteJsonResponse(w, map[string]interface{}{
			"token": token,
		}, http.StatusCreated)
	}
}

func WriteJsonError(w http.ResponseWriter, err errors.CustomError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.GetCode())
	json.NewEncoder(w).Encode(map[string]string{
		"error": err.GetMsg() + " Op:" + err.GetOp(),
	})
}

func WriteJsonResponse(w http.ResponseWriter, message map[string]interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(message)
}
