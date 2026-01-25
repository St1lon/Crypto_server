package main

import (
	//"cryptoserver/domain"
	"cryptoserver/handler"
	"cryptoserver/middleware"
	"cryptoserver/repository"
	"log"
	"net/http"
)
func main(){
	mux := http.NewServeMux()
	userRepo := repository.NewMemoryUserRepository()
	handlerWithAuth := middleware.AuthMiddleware(handler.HandlerAuth(userRepo))
	mux.HandleFunc("POST /auth/register", handler.HandlerRegister(userRepo))
	mux.Handle("/auth/login", handlerWithAuth)
    log.Println("Server starting on :8080")
    err := http.ListenAndServe(":8080", mux)
    if err != nil {
        log.Fatal(err)
    }
}