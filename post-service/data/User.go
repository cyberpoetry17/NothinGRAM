package data

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct{
	IDU      uuid.UUID `gorm:"column:id;PRIMARY_KEY"      json:"id"`
	Name     string `gorm:"column:name"      json:"name"`
	Surname  string `gorm:"column:surname"      json:"surname"`
	Username string `gorm:"column:username"      json:"username"`
	Private  bool   `gorm:"column:private" json:"private"`
}

func (user *User) BeforeCreate(scope *gorm.DB) error {
	user.IDU = uuid.New()
	return nil
}