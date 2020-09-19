package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

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
	if len(url.Password) > 0 {
		hashPass, err := lib.HashPassword(url.Password)
		if err != nil {
			log.Log().Errorln(err)
			resp.Title = "Something went wrong"
			resp.Status = http.StatusInternalServerError
			resp.Errors = err
			resp.SendResponse(w, r)
			return
		}
		url.Password = hashPass
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
	resp.Data = map[string]interface{}{
		"id":          url.ID,
		"isProtected": len(url.Password) > 0,
		"link":        url.Link,
	}
	resp.Status = http.StatusOK
	resp.SendResponse(w, r)
}

type URLResponse struct {
	ID          string `json:"id"`
	Link        string `json:"link,omitempty"`
	IsProtected bool   `json:"isProtected"`
}

// GetShortURL ...
func GetShortURL(w http.ResponseWriter, r *http.Request) {
	resp := core.Response()
	vars := mux.Vars(r)
	id := vars["id"]
	db := app.DB()
	urlRepo := data.NewURLRepository()
	url, err := urlRepo.GetURLByID(db, id)

	if err != nil {
		if isNotFound := lib.IsRecordNotFoundError(err); isNotFound {
			resp.Title = "Invalid short URL"
			resp.Status = http.StatusNotFound
			resp.SendResponse(w, r)
			return
		}
		log.Log().Errorln(err)
		resp.Title = "Database query failed"
		resp.Status = http.StatusInternalServerError
		resp.Errors = err
		resp.SendResponse(w, r)
		return
	}
	isPrivate := len(url.Password) > 0
	urlResp := URLResponse{ID: url.ID, IsProtected: isPrivate}
	if !isPrivate {
		urlResp.Link = url.Link
	}
	resp.Data = urlResp
	resp.Status = http.StatusOK
	resp.SendResponse(w, r)
}

// MatchURLPassword will match URL password and send id
func MatchURLPassword(w http.ResponseWriter, r *http.Request) {
	ru := models.URL{}
	resp := core.Response()
	if err := json.NewDecoder(r.Body).Decode(&ru); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db := app.DB()
	urlRepo := data.NewURLRepository()
	url, err := urlRepo.GetURLByID(db, ru.ID)

	if err != nil {
		if isNotFound := lib.IsRecordNotFoundError(err); isNotFound {
			resp.Title = "Invalid short URL"
			resp.Status = http.StatusNotFound
			resp.SendResponse(w, r)
			return
		}
		log.Log().Errorln(err)
		resp.Title = "Database query failed"
		resp.Status = http.StatusInternalServerError
		resp.Errors = err
		resp.SendResponse(w, r)
		return
	}
	if err := lib.CheckPassword(url.Password, ru.Password); err != nil {
		log.Log().Errorln(err)
		resp.Title = "Password not match"
		resp.Status = http.StatusUnauthorized
		resp.Errors = err
		resp.SendResponse(w, r)
		return
	}
	isPrivate := len(url.Password) > 0
	resp.Data = URLResponse{ID: url.ID, IsProtected: isPrivate, Link: url.Link}
	resp.Status = http.StatusOK
	resp.SendResponse(w, r)

}
