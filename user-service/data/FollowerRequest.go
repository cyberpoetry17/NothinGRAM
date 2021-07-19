package data

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FollowerRequest struct {
	ID         uuid.UUID `gorm:"primary_key,column:id"      json:"id"`
	IDFollower uuid.UUID `gorm:"column:idfollower"          json:"idfollower"`
	IDFollowed     uuid.UUID `gorm:"column:idfollowed"              json:"idfollowed"`
}

func (followerRequest *FollowerRequest) BeforeCreate(scope *gorm.DB) error {
	followerRequest.ID = uuid.New()
	return nil
}