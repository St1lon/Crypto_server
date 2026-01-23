package handler

import (
	"cryptoserver/domain"
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
    Password string `json:"password"`  // Приходит password, не hash!
}

func HandlerRegister(userRepo repository.UserRepository) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST"{
		http.Error(w, NewErrWrongMethod(r.Method, http.StatusMethodNotAllowed, "register user").Error(), http.StatusMethodNotAllowed)
		log.Println("wrong method:", r.Method)
		return
	}

	if ct := r.Header.Get("Content-Type"); ct != "application/json"{
		http.Error(w, NewErrWrongCT(ct, http.StatusUnsupportedMediaType, "register user").Error(), http.StatusUnsupportedMediaType)
		log.Println("unsupported content type:", ct)
		return
	}
	var user_request RegisterRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user_request);
	if  err != nil{
		http.Error(w, NewErrInvalidJSON(err.Error(), http.StatusBadRequest, "register user").Error(), http.StatusBadRequest)
		wrappedErr := fmt.Errorf("failed to decode JSON body: %w", err)
		log.Println(wrappedErr)
		return
	}
	if user_request.Username == ""{
		WriteJsonError(w, "Username is required", http.StatusBadRequest)
		return
	}
	if user_request.Password == ""{
		WriteJsonError(w, "Password is required", http.StatusBadRequest)
		return
	} 
	
	var user domain.User
	_, err := userRepo.GetByUsername(user_request.Username)
	if err == nil {
    WriteJsonError(w, "user already exists", http.StatusConflict)
    return
}
	user.Username = user_request.Username
	hash, err := bcrypt.GenerateFromPassword([]byte(user_request.Password), bcrypt.DefaultCost)
	if err != nil{
		WriteJsonError(w, "Fail to hasp password", http.StatusInternalServerError)
		return
	}
	user.PasswordHash = string(hash)
	err = userRepo.Create(&user)
	if err != nil{
		WriteJsonError(w, "Fail to create user", http.StatusConflict)
		return
	}
	
	token, err := middleware.GenerateToken(&user)
	if err != nil{
		WriteJsonError(w, "Error to generate JWT token", http.StatusInternalServerError)
	}
	WriteJsonResponse(w, map[string]interface{}{
		"token" : token,
	}, http.StatusCreated)
	}
}

func WriteJsonError(w http.ResponseWriter, err ErrCustom){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Code)
	json.NewEncoder(w).Encode(map[string]string{
		"error" : err.Msg + "Op:" + err.Op,
	})
}

func WriteJsonResponse(w http.ResponseWriter, message map[string]interface{}, code int){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(message)
}