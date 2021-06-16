package data

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type StoryType int
const(
	media MediaType = iota
	post
)

type Story struct{
	ID uuid.UUID `gorm:"column:id;PRIMARY_KEY"      json:"id"`
	Time time.Time `gorm:"column:Time;not null"      json:"Time"`
	Post Post `gorm:"foreignkey:ID"   json:"Post"`
	Media Media `gorm:"foreignkey:ID"   json:"Media"`
}

func (story *Story) BeforeCreate(scope *gorm.DB) error {
	story.ID = uuid.New()
	return nil

}
