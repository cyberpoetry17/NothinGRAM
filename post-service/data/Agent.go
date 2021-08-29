package data

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type  Agent struct {
	ID                   uuid.UUID       `json:"-"`
	Name                 string          `gorm:"column:name"   json:"name"`
	Surname              string          `gorm:"column:surname"  json:"surname"`
	Email                string          `gorm:"column:email"   json:"email"`
	Username             string          `gorm:"column:username" json:"username"`
	Password             string          `gorm:"column:password" json:"password"`
	DateOfBirth          time.Time       `gorm:"column:date" json:"date"`
	Gender               Gender          `gorm:"column:gender" json:"gender"`
	PhoneNumber          string          `gorm:"column:phone"  json:"phone"`
	Biography            string          `gorm:"column:bio"  json:"bio"`
	Website              string          `gorm:"column:web" json:"web"`
	Verified             bool            `gorm:"column:verify"  json:"verify"`
	Private              bool            `gorm:"column:private" json:"private"`
	Taggable             bool            `gorm:"column:taggable"  json:"taggable"`
	ReceiveNotifications bool            `gorm:"column:notif" json:"notifications"`
	Link              string          `gorm:"column:link"  json:"link"`

}

func (agent *Agent) BeforeCreate(scope *gorm.DB) error {
	agent.ID = uuid.New()
	return nil
}

type Gender int

const (
	Male Gender = iota + 1
	Female
)