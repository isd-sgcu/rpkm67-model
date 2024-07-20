package model

type Count struct {
	Base
	Name string `json:"name" gorm:"index"`
}
