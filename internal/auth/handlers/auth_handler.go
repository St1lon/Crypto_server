package handlers

import (
	"cryptoserver/errors"
	"cryptoserver/internal/repository"
	"cryptoserver/internal/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	//"github.com/golang-jwt/jwt/v5"
	//"cryptoserver/middleware"
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
			log.Printf("%s : %v", customErr, err)
			return
		}
		authService := service.NewAuthService(userRepo)
		token, customErr := authService.Register(user_request.Username, user_request.Password)
		if customErr != nil {
			WriteJsonError(w, customErr)
			log.Printf("%s", customErr.Error())
			return
		}
		WriteJsonResponse(w, map[string]string{
			"token": token,
		}, http.StatusCreated)
		log.Println("user registered:", user_request.Username)
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
