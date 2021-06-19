package services

import (
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/repository"
)

type BlockedService struct {
	Repo *repository.BlockedRepo
}

func (service *BlockedService) CreateBlockedUser(blockedUser *data.Blocked) error {
	return service.Repo.CreateBlocked(blockedUser)
}

func (service *BlockedService) GetAllBlockedUsers(userID string) ([]data.Blocked, error) {

	blockedUsers, error := service.Repo.GetAllBlockedUsersByID(userID)
	if error != nil {
		return nil, error
	}
	return blockedUsers, nil
}

func (service *BlockedService) RemoveBlockedUser(blockedUser *data.Blocked) error {
	return service.Repo.RemoveBlocked(blockedUser)
}
