package services

import (
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/repository"
)

type FollowerService struct {
	Repo *repository.FollowerRepo
}

func (service *FollowerService) FollowUser(follower *data.Follower) error {
	return service.Repo.FollowUser(follower)
}

func (service *FollowerService) UnfollowUser(follower *data.Follower) error {
	return service.Repo.UnfollowUser(follower)
}
