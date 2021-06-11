package services

import (
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/repository"
	"github.com/google/uuid"
)

type UserService struct {
	Repo *repository.UserRepo
}

//verovatno treba da vrati neku vrednost
func (service *UserService) CreateUser(user *data.User2) error {
	service.Repo.CreateUser(user)
	return nil
}

func (service *UserService) UserExists(userId string) (bool, error) {
	id, err := uuid.Parse(userId)
	if err != nil {
		print(err)
		return false, err
	}
	exists := service.Repo.UserExists(id)
	return exists, nil
}

func (service *UserService) UserExistsByEmail(email string) (bool, error) {
	exists := service.Repo.UserExistsByEmail(email)
	return exists, nil
}

func (service *UserService) UserExistsByUsername(username string) (bool, error) {
	exists := service.Repo.UserExistsByEmail(username)
	return exists, nil
}
