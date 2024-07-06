package model

import (
	"github.com/google/uuid"
	"github.com/isd-sgcu/rpkm67-model/utils"
	"gorm.io/gorm"
)

type Group struct {
	Base
	LeaderID    *uuid.UUID   `json:"leader_id"`
	Token       string       `json:"token" gorm:"index:, unique"`
	Members     []*User      `json:"members"`
	Selections  []*Selection `json:"selections"`
	IsConfirmed bool         `json:"is_confirmed"`
}

func (m *Group) BeforeCreate(_ *gorm.DB) error {
	m.Token = utils.GenGroupToken(m.LeaderID)
	m.ID = uuid.New()
	return nil
}
