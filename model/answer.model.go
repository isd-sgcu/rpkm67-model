package model

type Answer struct {
	Base
	ActivityID string `json:"activity_id" gorm:"index"`
	Text       string `json:"text"`
}
