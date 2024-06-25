package model

import "github.com/isd-sgcu/rpkm67-model/constant"

type User struct {
	Base
	Email     string        `json:"email" gorm:"tinytext;unique"`
	Password  string        `json:"password" gorm:"tinytext"`
	Firstname string        `json:"firstname" gorm:"tinytext"`
	Lastname  string        `json:"lastname" gorm:"tinytext"`
	Tel       string        `json:"tel" gorm:"tinytext"`
	Role      constant.Role `json:"role" gorm:"tinytext"`
}
