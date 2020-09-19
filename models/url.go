package models

// URL schema
type URL struct {
	ID       string `json:"id" gorm:"column:id;primary_key"`
	Link     string `json:"link" gorm:"column:link;not_null"`
	Password string `json:"password" gorm:"column:password"`
}

// TableName  return URL table name
func (u *URL) TableName() string {
	return "Urls"
}
