package data

import (
	"github.com/google/uuid"
)

type Comment struct {
	ID      uuid.UUID `gorm:"column:id"      json:"id"`
	Comment string    `gorm:"column:Comment"   json:"desciption"`
	UserId  string    `gorm:"column:UserId"   json:"UserId"`
}
