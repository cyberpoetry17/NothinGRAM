package repository

import (
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
)

type IUserRepository interface {
	GetUser(uint64) (*data.User2, error)
	GetAllUsers() (data.Users, error)
	SaveUser(*data.User2) (*data.User2, map[string]string)
	GetUserByEmailAndPassword(*data.User2) (*data.User2, map[string]string)
}
