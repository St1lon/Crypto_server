package handler

import (
	"net/http"
	"cryptoserver/domain"
	"cryptoserver/repository"
	"encoding/json"
)

type RegisterRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`  // Приходит password, не hash!
}

func handlerRegister(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST"{
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if ct := r.Header.Get("Content-Type"); ct != "application/json"{
		http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return
	}
	var user RegisterRequest

	decoder := json.NewDecoder(r.Body)
	if  err := decoder.Decode(&user); err != nil{
		WriteJsonError(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	if user.Username == ""{
		WriteJsonError(w, "Username is required", http.StatusBadRequest)
	}
	if user.Password == ""{
		WriteJsonError(w, "Password is required", http.StatusBadRequest)
	} 





}

func WriteJsonError(w http.ResponseWriter, message string, code int){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{
		"error" : message,
	})
}