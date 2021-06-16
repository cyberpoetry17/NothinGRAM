package data

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)
type MediaType int
const(
	picture MediaType = iota
	video
)

type Media struct{
	ID uuid.UUID `gorm:"column:id;PRIMARY_KEY"      json:"id"`
	Link string `gorm:"column:Link;not null"      json:"Link"`
	Type MediaType `gorm:"column:Type;not null"      json:"Type"`
	PostId  uuid.UUID `gorm:"column:PostId;not null"   json:"PostId"`
}

func (media *Media) BeforeCreate(scope *gorm.DB) error {
	media.ID = uuid.New()
	return nil

}