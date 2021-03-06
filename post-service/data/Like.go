package data

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Like struct{
	IDL uuid.UUID `gorm:"column:id;PRIMARY_KEY"      json:"id"`
	UserId string `gorm:"column:userid;not null"      json:"userid"`
	PostId uuid.UUID `gorm:"column:postid;not null"      json:"postid"`
}

func (like *Like) BeforeCreate(scope *gorm.DB) error {
	like.IDL = uuid.New()
	return nil
}