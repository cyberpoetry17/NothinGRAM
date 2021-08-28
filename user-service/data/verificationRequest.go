package data

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VerificationRequest struct {
	ID            uuid.UUID     `gorm:"primary_key,column:id"  json:"id"`
	Name          string        `gorm:"column:name"            json:"name"`
	Surname       string        `gorm:"column:surname"         json:"surname"`
	Username      string        `gorm:"column:username"        json:"username"`
	Category      Category      `gorm:"column:category"        json:"category"`
	RequestStatus RequestStatus `gorm:"column:status"          json:"status"`
}

func (verificationRequest *VerificationRequest) BeforeCreate(scope *gorm.DB) error {
	verificationRequest.ID = uuid.New()
	return nil
}

type Category int

const (
	Influencer Category = iota + 1
	Sports
	NewsOrMedia
	Business
	Brand
	Organization
)

func (d Category) EnumIndex() int {
	return int(d)
}

type RequestStatus int

const (
	Accepted RequestStatus = iota + 1
	Rejected
	Waitlisted
)
