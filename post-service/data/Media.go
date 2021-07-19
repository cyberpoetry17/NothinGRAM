package data

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)
type MediaType int
const(
	Picture MediaType = iota
	Video
)

type Media struct{
	ID uuid.UUID `gorm:"column:id;PRIMARY_KEY"      json:"id"`
	Link string `gorm:"column:Link"      json:"Link"`
	Type MediaType `gorm:"column:Type"      json:"Type"`
	StoryId *uuid.UUID `gorm:"column:StoryId"   json:"StoryId"`
	PostId  *uuid.UUID `gorm:"column:PostId"   json:"PostId"`

}

func (media *Media) BeforeCreate(scope *gorm.DB) error {
	if(media.ID.String() == "00000000-0000-0000-0000-000000000000") {
		media.ID = uuid.New()
	}

	return nil

}