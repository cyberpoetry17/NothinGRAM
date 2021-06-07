package repository

import (
	"github.com/cyberpoetry17/UserAPI/data"
)

type UserRepository interface {
	SaveUser(*data.User2) (*data.User2, map[string]string)
	GetUser(uint64) (*data.User2, error)
	GetUsers() ([]data.User2, error)
	GetUserByEmailAndPassword(*data.User2) (*data.User2, map[string]string)
}
