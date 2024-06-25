package model

type Stamp struct {
	Base
	PointA int    `json:"point_a"`
	PointB int    `json:"point_b"`
	PointC int    `json:"point_c"`
	PointD int    `json:"point_d"`
	Stamp  string `json:"stamp" gorm:"tinytext"` // e.g. 01000100010 = 2nd workshop + 1st landmark + 1st institute
	UserID int
	User   User
}
