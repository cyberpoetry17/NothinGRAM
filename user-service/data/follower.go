package data

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Follower struct {
	ID         uuid.UUID `gorm:"primary_key,column:id"      json:"id"`
	IDFollower uuid.UUID `gorm:"column:idfollower"          json:"idfollower"`
	IDUser     uuid.UUID `gorm:"column:iduser"              json:"iduser"`
}

func (follower *Follower) BeforeCreate(scope *gorm.DB) error {
	follower.ID = uuid.New()
	return nil
}
