package model

import (
	"github.com/google/uuid"
	"github.com/isd-sgcu/rpkm67-model/constant"
)

type User struct {
	Base
	Email       string        `json:"email" gorm:"tinytext;unique"`
	Nickname    string        `json:"nickname" gorm:"tinytext"`
	Title       string        `json:"title" gorm:"tinytext"`
	Firstname   string        `json:"firstname" gorm:"tinytext"`
	Lastname    string        `json:"lastname" gorm:"tinytext"`
	Year        int           `json:"year" gorm:"type:smallint"`
	Faculty     string        `json:"faculty" gorm:"tinytext"`
	Tel         string        `json:"tel" gorm:"tinytext"`
	ParentTel   string        `json:"parent_tel" gorm:"tinytext"`
	Parent      string        `json:"parent" gorm:"tinytext"`
	FoodAllergy string        `json:"food_allergy" gorm:"tinytext"`
	DrugAllergy string        `json:"drug_allergy" gorm:"tinytext"`
	Illness     string        `json:"illness" gorm:"tinytext"`
	Role        constant.Role `json:"role" gorm:"tinytext"`
	PhotoKey    string        `json:"photo_key" gorm:"tinytext"`
	PhotoUrl    string        `json:"photo_url" gorm:"tinytext"`
	Baan        string        `json:"baan" gorm:"tinytext"`
	ReceiveGift int           `json:"receive_gift" gorm:"tinyint"`
	GroupID     *uuid.UUID    `json:"group_id"`
	CheckIns    []*CheckIn    `json:"check_ins"`
}
