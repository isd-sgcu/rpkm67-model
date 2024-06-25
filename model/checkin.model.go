package model

type Checkin struct {
	Base
	Event  string `json:"event" gorm:"tinytext"`
	Email  string `json:"email" gorm:"tinytext"`
	UserID uint   `json:"user_id"`
}
