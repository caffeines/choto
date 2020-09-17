package api

import (
	"encoding/json"
	"net/http"

	"github.com/caffeines/choto/lib"
	"github.com/caffeines/choto/log"

	"github.com/caffeines/choto/app"
	"github.com/caffeines/choto/data"

	"github.com/caffeines/choto/models"

	"github.com/caffeines/choto/core"
)

func createURL(url models.URL) error {
	db := app.DB().Begin()
	urlRepo := data.NewURLRepository()
	err := urlRepo.CreateURL(db, &url)
	if err != nil {
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
	url := models.URL{}

	if err := json.NewDecoder(r.Body).Decode(&url); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var isUserDefinedID bool
	if len(url.ID) == 0 {
		url.ID = lib.RandStringRunes(6)
	} else {
		isUserDefinedID = true
	}
	var err error
	for {
		err = createURL(url)
		if _, isIt := lib.IsDuplicateKeyError(err); isIt {
			if isUserDefinedID {
				resp.Title = "This id already taken"
				resp.Status = http.StatusConflict
				resp.Errors = err
				resp.SendResponse(w, r)
				return
			}
			url.ID = lib.RandStringRunes(6)
		} else {
			break
		}
	}
	if err != nil {
		log.Log().Errorln(err)
		resp.Title = "Database query failed"
		resp.Status = http.StatusInternalServerError
		resp.Errors = err
		resp.SendResponse(w, r)
		return
	}
	resp.Data = url
	resp.Status = http.StatusOK
	resp.SendResponse(w, r)
}
