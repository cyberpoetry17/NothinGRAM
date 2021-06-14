package data

import (
	"github.com/google/uuid"
)

type Post struct{
	ID                   uuid.UUID `gorm:"primary_key,column:id"      json:"id"` //mozda ovo ne treba?
	Description			 string `gorm:"column:description"   json:"desciption"`
	PicturePath			 string `gorm:"column:picpath"   json:"picpath"`
	Likes				 int64	`gorm:"column:likes"   json:"likes"`
	Dislikes			 int64	`gorm:"column:dislikes"   json:"dislikes"`
	UserID				 uuid.UUID `gorm:"column:userid"   json:"userid"`
	Timestamp			 string `gorm:"column:timestamp"   json:"timestamp"`
	Tags 				[]Tag `gorm:"many2many:posts_tags;"   json:"Tags"`
	Comments			[]Comment `gorm:"foreignkey:PostId"   json:"Comments"`
}