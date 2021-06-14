package data

import (
	"github.com/google/uuid"
)

type Tag struct {
	ID      uuid.UUID `gorm:"column:id"      json:"id"`
	TagName string    `gorm:"column:TagName"   json:"TagName"`
}
