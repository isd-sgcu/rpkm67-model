package model

type CheckIn struct {
	Base
	Event  string `json:"event" gorm:"tinytext"`
	Email  string `json:"email" gorm:"tinytext"`
	UserID uint   `json:"user_id"`
}
