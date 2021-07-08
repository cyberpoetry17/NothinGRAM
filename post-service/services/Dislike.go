package services

import (
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/repository"
)

type DislikeService struct {
	Repo *repository.DislikeRepo
}

func (service *DislikeService) CreateDislike (dislike *data.Dislike) error{
	return service.Repo.CreateDislike(dislike)
}

func (service *DislikeService) GetAllDislikesForPost (postId string) []data.Dislike{
	return service.Repo.GetAllDislikesForPost(postId)
}

func (service *DislikeService) RemoveDislike (dislike *data.Dislike) error{
	return service.Repo.RemoveDislike(dislike)
}

func (service *DislikeService) CheckIfUserDislikedPost (dislike *data.Dislike) bool{
	return service.Repo.CheckIfUserDislikedPost(dislike)
}