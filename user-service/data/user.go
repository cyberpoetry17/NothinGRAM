package data

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User2 struct {
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
	Role                 Role            `gorm:"column:role"  json:"role"`
	Verified             bool            `gorm:"column:verify"  json:"verify"`
	Private              bool            `gorm:"column:private" json:"private"`
	Taggable             bool            `gorm:"column:taggable"  json:"taggable"`
	ReceiveNotifications bool            `gorm:"column:notif" json:"notifications"`
	Followers            []Follower      `gorm:"foreignkey:IDUser" json:"followers"`
	Following            []Follower      `gorm:"foreignkey:IDFollower" json:"following"`
	MutedUsers           []Muted         `gorm:"foreignkey:UserID"   json:"mutedUsers"`
	BlockedUsers         []Blocked       `gorm:"foreignkey:UserID"   json:"blockedUsers"`
	CloseFollowers       []CloseFollower `gorm:"foreignkey:IDUser" json:"closeFollowers"`
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

func (u *User2) IsCredentialsVerified(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User2) SetPassword(password string) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	u.Password = string(hashedPassword)
}

func (user *User2) BeforeCreate(scope *gorm.DB) error {
	user.ID = uuid.New()
	return nil
}

func (user *User2) AfterDelete(scope *gorm.DB) error {
	return scope.Model(&Follower{}).Where("iduser = ?", user.ID).Unscoped().Delete(&Follower{}).Error
}
