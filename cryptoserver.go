package main

import (
	//"cryptoserver/domain"
	"cryptoserver/internal/controller"
	"cryptoserver/internal/adapter"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	userRepo := adapter.NewMemoryUserRepository()
	mux.HandleFunc("POST /auth/register", controller.HandlerRegister(userRepo))
	mux.HandleFunc("POST /auth/login", controller.HandlerLogin(userRepo))
	log.Println("Server starting on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
