package server

import(
	"net/http"
)
type server struct{
	router  *http.Server
}