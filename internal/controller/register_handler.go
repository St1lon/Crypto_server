package controller

import (
	"cryptoserver/internal/domain"
	"cryptoserver/internal/service"
	"encoding/json"
	"log"
	"net/http"
)

func HandlerRegister(userRepo service.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			customErr := domain.NewErrWrongMethod("wrong method: "+r.Method, http.StatusMethodNotAllowed, "register user")
			WriteJsonError(w, customErr)
			log.Println(customErr)
			return
		}

		if ct := r.Header.Get("Content-Type"); ct != "application/json" {
			customErr := domain.NewErrWrongCT("Content-Type must be application/json", http.StatusUnsupportedMediaType, "register user")
			WriteJsonError(w, customErr)
			log.Println("unsupported content type:", customErr)
			return
		}
		var user_request RegisterRequest
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&user_request)
		if err != nil {
			customErr := domain.NewErrInvalidJSON(err.Error(), http.StatusBadRequest, "register user")
			WriteJsonError(w, customErr)
			log.Printf("%s : %v", customErr, err)
			return
		}
		authService := service.NewAuthService(userRepo)
		token, customErr := authService.RegisterUser(user_request.Username, user_request.Password)
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
