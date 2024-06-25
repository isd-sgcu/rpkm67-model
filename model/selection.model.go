package model

import (
	"time"

	"github.com/google/uuid"
)

type Selection struct {
	GroupID   *uuid.UUID `json:"group_id" gorm:"index"`
	BaanID    *uuid.UUID `json:"baan_id" gorm:"index"`
	Order     int        `json:"order" gorm:"type:smallint"`
	CreatedAt time.Time  `json:"created_at" gorm:"type:timestamp;autoCreateTime:nano"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"type:timestamp;autoUptimestamp:nano"`
}
