package services

import (
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/repository"
)

type FollowerRequestService struct {
	Repo *repository.FollowerRequestRepo
}

func (service *FollowerRequestService) CreateFollowRequest(followerRequest *data.FollowerRequest) error {
	return service.Repo.CreateFollowRequest(followerRequest)
}

func (service *FollowerRequestService) GetAllRequests(id string) []data.FollowerRequest{
	return service.Repo.GetAllForUser(id)
}

func (service *FollowerRequestService) RemoveRequest(id string) error {
	return service.Repo.RemoveRequest(id)
}