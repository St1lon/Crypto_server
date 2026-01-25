package main

import (
	//"cryptoserver/domain"
	"cryptoserver/internal/auth/handlers"
	"cryptoserver/internal/repository"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	userRepo := repository.NewMemoryUserRepository()
	mux.HandleFunc("POST /auth/register", handlers.HandlerRegister(userRepo))
	mux.HandleFunc("POST /auth/login", handlers.HandlerAuth(userRepo))

	log.Println("Server starting on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
