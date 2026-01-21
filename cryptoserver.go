package main

import(
	"net/http"
)
func main(){
	mux := http.NewServeMux()

	mux.HandleFunc("POST /auth/register", )
	mux.HandleFunc("POST /auth/login", )

}