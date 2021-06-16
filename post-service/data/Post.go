package data

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Post struct{
	ID                   uuid.UUID `gorm:"primary_key,column:id"      json:"id"` //mozda ovo ne treba?
	Description			 string `gorm:"column:description"   json:"desciption"`
	PicturePath			 string `gorm:"column:picpath"   json:"picpath"`
	Likes				 []Like	`gorm:"foreignKey:PostId"   json:"likes"`
	Dislikes			 []Dislike	`gorm:"foreignKey:PostId"   json:"dislikes"`
	UserID				 uuid.UUID `gorm:"column:userid"   json:"userid"`
	Timestamp			 time.Time `gorm:"column:timestamp"   json:"timestamp"`
	Tags 				[]Tag `gorm:"many2many:posts_tags;"   json:"Tags"`
	Comments			[]Comment `gorm:"foreignkey:PostId"   json:"Comments"`
	Media				[]Media `gorm:"foreignkey:PostId"   json:"Media"`
}

func (post *Post) BeforeCreate(scope *gorm.DB) error {
	post.ID = uuid.New()
	return nil

}