package server

import (
	"net/http"

	"github.com/caffeines/choto/api"
	"github.com/gorilla/mux"
)

// GetRouter ...
func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", api.CreateShortURL).Methods(http.MethodPost)
	return r
}
