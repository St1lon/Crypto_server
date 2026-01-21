package main

import (
	//"cryptoserver/domain"
	"cryptoserver/handler"
	"cryptoserver/repository"
	"net/http"
	"log"
)
func main(){
	mux := http.NewServeMux()
	userRepo := repository.NewMemoryUserRepository()
	mux.HandleFunc("POST /auth/register", handler.HandlerRegister(userRepo))
	//mux.HandleFunc("POST /auth/login", )

    log.Println("Server starting on :8080")
    err := http.ListenAndServe(":8080", mux)
    if err != nil {
        log.Fatal(err)
    }
}