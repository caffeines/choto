package api

import (
	"net/http"

	"github.com/caffeines/choto/core"
)

// CreateShortURL ...
func CreateShortURL(w http.ResponseWriter, r *http.Request) {
	resp := core.Response()
	resp.Data = map[string]interface{}{
		"name": "Sadat",
	}
	resp.Status = http.StatusOK
	resp.SendResponse(w, r)
}
