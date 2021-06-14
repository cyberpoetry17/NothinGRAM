package data

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct{
	ID                   uuid.UUID `gorm:"column:id"      json:"id"` //mozda ovo ne treba?
	Description			 string `gorm:"column:description"   json:"description"`
	PicturePath			 string `gorm:"column:picpath"   json:"picpath"`
	Likes				 []Like	`gorm:"foreignKey:PostId"   json:"likes"`
	Dislikes			 []Dislike	`gorm:"foreignKey:PostId"   json:"dislikes"`
	UserID				 uuid.UUID `gorm:"column:userid"   json:"userid"`
	Timestamp			 string `gorm:"column:timestamp"   json:"timestamp"`
}

func (post *Post) BeforeCreate(scope *gorm.DB) error {
	post.ID = uuid.New()
	return nil
}