package data

import (
	"github.com/caffeines/choto/models"
	"github.com/jinzhu/gorm"
)

type URLRepositoryImpl struct{}

var urlRepo URLRepository

// NewURLRepository will return URLRepository
func NewURLRepository() URLRepository {
	if urlRepo == nil {
		urlRepo = &URLRepositoryImpl{}
	}
	return urlRepo
}

// CreateURL create new URL in DB
func (u *URLRepositoryImpl) CreateURL(db *gorm.DB, url *models.URL) error {
	if err := db.Model(url).Create(url).Error; err != nil {
		return err
	}
	return nil
}

// GetURLByID return URL from DB
func (u *URLRepositoryImpl) GetURLByID(db *gorm.DB, id string) (*models.URL, error) {
	url := models.URL{}
	if err := db.Model(&url).Where("id = ?", id).First(&url).Error; err != nil {
		return nil, err
	}
	return &url, nil
}
