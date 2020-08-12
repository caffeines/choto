package models

// Url schema
type Url struct {
	ID       string `json:"id" gorm:"column:id;primary_key"`
	Link     string `json:"link" gorm:"column:link;not_null"`
	Password string `json:"password" gorm:"column:password"`
}

func (u *Url) TableName() string {
	return "Urls"
}
