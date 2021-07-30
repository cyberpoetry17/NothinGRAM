package data

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type StoryType int
const(
	MediaT StoryType = iota
	PostT
)

type Story struct{
	IdStory uuid.UUID `gorm:"column:IdStory;PRIMARY_KEY"      json:"IdStory"`
	Time time.Time `gorm:"column:Time;not null"      json:"Time"`
	MediaID uuid.UUID `gorm:"column:MediaID"   json:"MediaID"`
	PostID uuid.UUID `gorm:"column:PostID"   json:"PostID"`
	Type StoryType `gorm:"column:Type"  json:"Type"`
	IsActive bool `gorm:"column:IsActive"  json:"IsActive"`
	UserId uuid.UUID `gorm:"column:UserId"  json:"UserId"`
	IsOnlyForCloseFriends bool `gorm:"column:IsOnlyForCloseFriends"  json:"IsOnlyForCloseFriends"`
}

func (story *Story) BeforeCreate(scope *gorm.DB) error {
	if(story.IdStory.String() == "00000000-0000-0000-0000-000000000000") {
		story.IdStory = uuid.New()
	}
	return nil

}
