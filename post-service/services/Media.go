package services

import (
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/repository"
	"github.com/google/uuid"
)

type MediaService struct {
	Repo *repository.MediaRepo
}

func (services *MediaService) CreateMedia(media *data.Media) error {
	return services.Repo.CreateMedia(media)
}

func (services *MediaService) RemoveMedia(media *data.Media) error {
	return services.Repo.RemoveMedia(media)
}

func (services *MediaService) EditMedia(media *data.Media) error {
	return services.Repo.EditMedia(media)
}

func (services *MediaService) GetMediaForPost(postID uuid.UUID) (error,[]data.Media) {
	return services.Repo.GetMediaForPost(postID)
}