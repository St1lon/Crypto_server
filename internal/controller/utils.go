package controller

import (
	"cryptoserver/internal/domain"
	"encoding/json"
	"fmt"
	"net/http"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func WriteJsonError(w http.ResponseWriter, err domain.CustomError) {
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
