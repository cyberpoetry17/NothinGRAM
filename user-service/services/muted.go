package services

import (
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/repository"
)

type MutedService struct {
	Repo *repository.MutedRepo
}

func (service *MutedService) CreateMutedUser(mutedUser *data.Muted) error {
	return service.Repo.CreateMuted(mutedUser)
}

func (service *MutedService) RemoveMutedUser(mutedUser *data.Muted) error {
	return service.Repo.RemoveMuted(mutedUser)
}

func (service *MutedService) GetAllMutedUsers(userID string) ([]data.Muted, error) {

	mutedUsers, error := service.Repo.GetAllMutedUsersByID(userID)
	if error != nil {
		return nil, error
	}
	return mutedUsers, nil
}
