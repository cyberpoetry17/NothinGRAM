package data

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Location struct{
	IDLoc   uuid.UUID `gorm:"column:id;PRIMARY_KEY"      json:"id"`
	Country string `gorm:"column:country"      json:"country"`
	City    string `gorm:"column:city"      json:"city"`
	Address string `gorm:"column:address"      json:"address"`
	PostId  uuid.UUID `gorm:"column:postid"      json:"postid"`

}

func (location *Location) BeforeCreate(scope *gorm.DB) error {
	location.IDLoc = uuid.New()
	return nil
}