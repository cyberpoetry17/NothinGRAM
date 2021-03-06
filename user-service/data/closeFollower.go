package data

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CloseFollower struct {
	ID              uuid.UUID `gorm:"primary_key,column:id"      json:"id"`
	IDCloseFollower uuid.UUID `gorm:"column:idfollower"          json:"idclosefollower"`
	IDUser          uuid.UUID `gorm:"column:iduser"              json:"iduser"`
}

func (closeFollower *CloseFollower) BeforeCreate(scope *gorm.DB) error {
	closeFollower.ID = uuid.New()
	return nil
}
