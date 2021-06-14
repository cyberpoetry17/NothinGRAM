package data

import (
	"github.com/google/uuid"
)

type Comment struct {
	ID      uuid.UUID `gorm:"primary_key,column:id"      json:"id"`
	Comment string    `gorm:"column:Comment"   json:"desciption"`
	UserId  string    `gorm:"column:UserId"   json:"UserId"`
	PostId  uuid.UUID `gorm:"column:UserId"   json:"PostId"`
}
