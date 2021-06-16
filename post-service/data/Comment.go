package data

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	ID      uuid.UUID `gorm:"primary_key,column:id"      json:"id"`
	Comment string    `gorm:"column:Comment"   json:"Comment"`
	UserId  uuid.UUID   `gorm:"column:UserId;not null"   json:"UserId"`
	PostId  uuid.UUID `gorm:"column:PostId;not null"   json:"PostId"`
}

func (comment *Comment) BeforeCreate(scope *gorm.DB) error {
	comment.ID = uuid.New()
	return nil
}
