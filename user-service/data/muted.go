package data

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Muted struct {
	ID      uuid.UUID `gorm:"column:id;PRIMARY_KEY"      json:"id"`
	UserID  uuid.UUID `gorm:"column:userID;not null"   json:"UserID"`
	MutedID uuid.UUID `gorm:"column:mutedID;not null"   json:"MutedID"`
}

func (muted *Muted) BeforeCreate(scope *gorm.DB) error {
	muted.ID = uuid.New()
	return nil
}
