package model

import (
	"os/user"

	"github.com/google/uuid"
	"github.com/isd-sgcu/rpkm67-model/utils"
	"gorm.io/gorm"
)

type Group struct {
	Base
	LeaderID   string       `json:"leader_id"`
	Token      string       `json:"token" gorm:"index:, unique"`
	Members    []*user.User `json:"members"`
	Selections []*Selection `json:"selections"`
}

func (m *Group) BeforeCreate(_ *gorm.DB) error {
	m.Token = utils.GenGroupToken(m.LeaderID)
	m.ID = uuid.New()
	return nil
}
