package data

import (
	"github.com/google/uuid"
)

type Post struct {
	ID      uuid.UUID `gorm:"column:id"      json:"id"`
	TagName string    `gorm:"column:TagName"   json:"TagName"`
}
