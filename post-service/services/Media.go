package services

import (
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/repository"
)

type MediaService struct {
	Repo *repository.MediaRepo
}

func (service *MediaService) CreateMedia(media *data.Media) error {
	return service.CreateMedia(media)
}

func (service *MediaService) RemoveMedia(media *data.Media) error {
	return service.RemoveMedia(media)
}

func (service *MediaService) EditMedia(media *data.Media) error {
	return service.EditMedia(media)
}