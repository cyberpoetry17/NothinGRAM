package services

import (
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/DTO"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/repository"
	"github.com/google/uuid"
)

type CloseFollowerService struct {
	Repo *repository.CloseFollowerRepository
}

func (service *CloseFollowerService) AddCloseFollower(follower *data.CloseFollower) error {
	return service.Repo.AddCloseFollower(follower)
}

func (service *CloseFollowerService) RemoveCloseFollower(follower *data.CloseFollower) error {
	return service.Repo.RemoveCloseFollower(follower)
}

func (service *CloseFollowerService) AddMultipleFollowers(newFollowerId []DTO.UserDTO, id uuid.UUID) error {
	for _, element := range newFollowerId {
		var closeFollower data.CloseFollower
		closeFollower.IDCloseFollower = element.UserId
		closeFollower.IDUser = id
		err := service.Repo.AddCloseFollower(&closeFollower)
		if err != nil {
			return err
		}
	}
	return nil
}

func (service *CloseFollowerService) RemoveManyByID(ids []DTO.UserDTO) error {
	for _, element := range ids {
		err := service.Repo.RemoveManyByID(element)
		if err != nil {
			return err
		}
	}
	return nil
}
