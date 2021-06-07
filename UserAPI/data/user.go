package data

import (
	"encoding/json"
	"io"

	"gorm.io/gorm"
)

// "encoding/json"
// "fmt"
// "io"
// "time"

type User2 struct {
	gorm.Model

	Name                 string `gorm:"column:name"   json:"name"`
	Surname              string `gorm:"column:surname"  json:"surname"`
	Email                string `gorm:"column:email"   json:"email"`
	Username             string `gorm:"column:username" json:"username"`
	Password             string `gorm:"column:password" json:"password"`
	DateOfBirth          string `gorm:"column:date" json:"date"`
	Gender               Gender `gorm:"column:gender" json:"gender"`
	PhoneNumber          string `gorm:"column:phone"  json:"phone"`
	Biography            string `gorm:"column:bio"  json:"bio"`
	Website              string `gorm:"column:web" json:"web"`
	Role                 Role   `gorm:"column:role"  json:"role"`
	Verified             bool   `gorm:"column:verify"  json:"verify"`
	Private              bool   `gorm:"column:private" json:"private"`
	Taggable             bool   `gorm:"column:taggable"  json:"taggable"`
	ReceiveNotifications bool   `gorm:"column:notif" json:"notifications"`
	//MutedId              []string `gorm:"one2many:muted" json:"muted"` //gorm:"many2many:article_tag"
	//BlockedId            []string `gorm:"column:blocked" json:"blocked"` *MultiString `gorm:"type:text[]"`
}
type Role int

type Users []*User2 //lista usera

const (
	Admin Role = iota + 1
	RegisteredUser
	Agent
)

func (d Role) EnumIndex() int {
	return int(d)
}

type Gender int

const (
	Male Gender = iota + 1
	Female
)

func (d Gender) EnumIndex() int {
	return int(d)
}

//dekoduje usera u JSON format za slanje putem RESTa,ako ne-> vraca gresku
func (user *Users) FromJSON(response io.Reader) error {
	err := json.NewDecoder(response)
	return err.Decode(user)
}

//prebacuje iz JSONa u GO
func (user *Users) ToJSON(write io.Writer) error {
	err := json.NewEncoder(write)
	return err.Encode(user)
}

// func GetAllUsers() Users {
// 	//return usersList //za sada
// 	return nil
// }
