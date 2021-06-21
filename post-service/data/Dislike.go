package data

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Dislike struct{
	IDD  uuid.UUID `gorm:"column:id;PRIMARY_KEY"      json:"id"`
	UserId string `gorm:"column:userid;not null"      json:"userid"`
	PostId uuid.UUID `gorm:"column:postid;not null"      json:"postid"`
}

func (dislike *Dislike) BeforeCreate(scope *gorm.DB) error {
	dislike.IDD = uuid.New()
	return nil
}