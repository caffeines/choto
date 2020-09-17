package core

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status int         `json:"-"`
	Title  string      `json:"title,omitempty"`
	Data   interface{} `json:"data,omitempty"`
	Errors error       `json:"errors,omitempty"`
}

// SendResponse ...
func (r Response) SendResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)
	if err := json.NewEncoder(w).Encode(r.Data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return nil
}
