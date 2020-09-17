package data

import (
	"github.com/caffeines/choto/models"
	"github.com/jinzhu/gorm"
)

type URLRepository interface {
	CreateURL(db *gorm.DB, url *models.URL) error
	GetURLByID(db *gorm.DB, id string) (*models.URL, error)
}
