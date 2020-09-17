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

func createURL(url models.URL) error {
	db := app.DB().Begin()
	urlRepo := data.NewURLRepository()
	err := urlRepo.CreateURL(db, &url)
	if err != nil {
		log.Log().Errorln(err)
		db.Rollback()
		return err
	}
	if err := db.Commit().Error; err != nil {
		return err
	}
	return nil
}

// CreateShortURL ...
func CreateShortURL(w http.ResponseWriter, r *http.Request) {
	resp := core.Response()
	url := models.URL{ID: lib.RandStringRunes(6)}

	if err := json.NewDecoder(r.Body).Decode(&url); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := createURL(url)
	if _, isIt := lib.IsDuplicateKeyError(err); isIt {
		url.ID = lib.RandStringRunes(6)
		err = createURL(url)
	}
	if err != nil {
		resp.Title = "Database query failed"
		resp.Status = http.StatusInternalServerError
		resp.Errors = err
		resp.Data = nil
		resp.SendResponse(w, r)
		return
	}
	resp.Data = url
	resp.Status = http.StatusOK
	resp.SendResponse(w, r)
}
