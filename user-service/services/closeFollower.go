package services

import (
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/repository"
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
