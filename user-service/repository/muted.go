package repository

import (
	"fmt"

	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"gorm.io/gorm"
)

type MutedRepo struct {
	Database *gorm.DB
}

func (repo MutedRepo) CreateLike(muted *data.Muted) error {
	result := repo.Database.Create(muted)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println(result.RowsAffected)
	return nil
}

func (repo MutedRepo) RemoveMuted(muted *data.Muted) error {
	return repo.Database.Delete(muted).Error
}
