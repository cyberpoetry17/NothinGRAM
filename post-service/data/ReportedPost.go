package data

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)
type ReportedPost struct{
	IDR uuid.UUID `gorm:"column:id;PRIMARY_KEY"      json:"id"`
	UserId string `gorm:"column:userid;not null"      json:"userid"`
	PostId uuid.UUID `gorm:"column:postid;not null"      json:"postid"`
}

func (report *ReportedPost) BeforeCreate(scope *gorm.DB) error {
	report.IDR = uuid.New()
	return nil
}