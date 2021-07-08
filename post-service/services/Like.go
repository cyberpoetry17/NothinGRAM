package services

import (
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/repository"
)

type LikeService struct {
	Repo *repository.LikeRepo
}

func (service *LikeService) CreateLike (like *data.Like) error{
	return service.Repo.CreateLike(like)
}

func (service *LikeService) GetAllLikesForPost (postId string) []data.Like{
	return service.Repo.GetAllLikesForPost(postId)
}

func (service *LikeService) RemoveLike (like *data.Like) error{
	return service.Repo.RemoveLike(like)
}

func (service *LikeService) CheckIfUserLikedPost (like *data.Like) bool{
	return service.Repo.CheckIfUserLikedPost(like)
}