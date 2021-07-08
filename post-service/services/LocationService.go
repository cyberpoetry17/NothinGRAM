package services

import (
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/repository"
	"github.com/google/uuid"
)

type LocationService struct {
	Repo *repository.LocationRepo
}

func (service *LocationService) CreateLocation (location *data.Location) (error,uuid.UUID){
	return service.Repo.CreateLocation(location)
}

func (service *LocationService) GetLocationForPost (postId string) *data.Location{
	return service.Repo.GetLocationForPost(postId)
}

func (service *LocationService) RemoveLocation (location *data.Location) error{
	return service.Repo.RemoveLocation(location)
}

func (service *LocationService) FilterPublicMaterialByLocations(locationid string) []data.Post{
	return service.Repo.FilterPublicMaterialByLocation(locationid)
}