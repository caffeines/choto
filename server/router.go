package server

import (
	"net/http"

	"github.com/caffeines/choto/api"
	"github.com/gorilla/mux"
)

// GetRouter ...
func GetRouter() *mux.Router {
	r := mux.NewRouter()
	postRoute := r.Methods(http.MethodPost).Subrouter()
	postRoute.HandleFunc("/hello", api.CreateShortURL)
	return r
}
