package data

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Dislike struct{
	IDD  uuid.UUID `gorm:"column:id"      json:"id"`
	UserId string `gorm:"column:userid"      json:"userid"`
	PostId uuid.UUID `gorm:"column:postid"      json:"postid"`
}

func (dislike *Dislike) BeforeCreate(scope *gorm.DB) error {
	dislike.IDD = uuid.New()
	return nil
}