package repository

import (
	"fmt"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"gorm.io/gorm"
)

type LocationRepo struct {
	Database *gorm.DB
}

func (repo LocationRepo) CreateLocation(location *data.Location) error {
	result := repo.Database.Create(location)
	if(result.Error != nil){
		return result.Error
	}
	fmt.Println(result.RowsAffected)
	return nil
}

func (repo LocationRepo) GetLocationForPost (postId string) *data.Location{
	var location data.Location
	repo.Database.Find(&location).Where("PostId = ?",postId)
	repo.Database.Preload("Posts",&location)
	return &location
}

func (repo LocationRepo) RemoveLocation (location *data.Location) error{
	return repo.Database.Delete(location).Error
}

func (repo *LocationRepo) GetAll() []data.Location{
	var locs []data.Location
	repo.Database.
		Preload("Posts").
		Find(&locs)
	return locs
}