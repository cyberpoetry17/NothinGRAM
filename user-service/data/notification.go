package data

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Notification struct {
	ID           uuid.UUID `gorm:"column:id;PRIMARY_KEY"      json:"id"`
	UserID       uuid.UUID `gorm:"column:userID;not null"   json:"userID"`
	Notification string    `gorm:"column:notification;not null"   json:"notification"`
}

func (notification *Notification) BeforeCreate(scope *gorm.DB) error {
	notification.ID = uuid.New()
	return nil
}
