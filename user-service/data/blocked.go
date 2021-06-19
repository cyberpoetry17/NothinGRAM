package data

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Blocked struct {
	ID        uuid.UUID `gorm:"column:id;PRIMARY_KEY"      json:"id"`
	UserID    uuid.UUID `gorm:"column:userID;not null"   json:"userID"`
	BlockedID uuid.UUID `gorm:"column:blockedID;not null"   json:"blockedID"`
}

func (blocked *Blocked) BeforeCreate(scope *gorm.DB) error {
	blocked.ID = uuid.New()
	return nil
}
