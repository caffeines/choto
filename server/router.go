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
	getRoute := r.Methods(http.MethodGet).Subrouter()
	postRoute.HandleFunc("/api", api.CreateShortURL)
	postRoute.HandleFunc("/api/match-password", api.MatchURLPassword)
	getRoute.HandleFunc("/api/{id}", api.GetShortURL)

	return r
}
