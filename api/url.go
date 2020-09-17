package api

import (
	"encoding/json"
	"net/http"

	"github.com/caffeines/choto/lib"

	"github.com/caffeines/choto/app"
	"github.com/caffeines/choto/data"
	"github.com/caffeines/choto/log"

	"github.com/caffeines/choto/models"

	"github.com/caffeines/choto/core"
)

// CreateShortURL ...
func CreateShortURL(w http.ResponseWriter, r *http.Request) {
	resp := core.Response()
	url := models.URL{}

	if err := json.NewDecoder(r.Body).Decode(&url); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db := app.DB().Begin()

	urlRepo := data.NewURLRepository()
	url.ID = lib.RandStringRunes(6)
	err := urlRepo.CreateURL(db, &url)
	if err != nil {
		db.Rollback()
		log.Log().Errorln(err)
		resp.Title = "Failed to create URL"
		resp.Status = http.StatusInternalServerError
		resp.Errors = err
		resp.SendResponse(w, r)
		return
	}
	resp.Data = url
	resp.Status = http.StatusOK
	resp.SendResponse(w, r)
}
