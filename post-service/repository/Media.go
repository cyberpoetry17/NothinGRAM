package repository

import (
	"fmt"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"gorm.io/gorm"
)

type MediaRepo struct {
	Database *gorm.DB
}

func (repo *MediaRepo) CreateMedia(media *data.Media) error {
	result := repo.Database.Create(media)
	if(result.Error != nil){
		return result.Error
	}
	fmt.Println(result.RowsAffected)
	return nil
}

func (repo *MediaRepo) RemoveMedia(media *data.Media) error {
	return repo.Database.Delete(media).Error
}

func (repo *MediaRepo) EditMedia(media *data.Media) error {
	return repo.Database.Save(media).Error
}