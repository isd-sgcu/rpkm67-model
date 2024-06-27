package model

type CheckIn struct {
	Base
	Event  string `json:"event" gorm:"tinytext"`
	Email  string `json:"email" gorm:"tinytext"`
	UserID string `json:"user_id"`
}
