package repository

import (
	"fmt"

	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MutedRepo struct {
	Database *gorm.DB
}

func (repo MutedRepo) CreateMuted(muted *data.Muted) error {
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

func (repo MutedRepo) GetAllMutedUsersByID(userID string) ([]data.Muted, error) {
	id, err := uuid.Parse(userID)
	if err != nil {
		print(err)
		return nil, err
	}
	var mutedUsers []data.Muted
	repo.Database.Find(&mutedUsers).Where("userID = ?", id)
	repo.Database.Preload("user2", &mutedUsers)
	return mutedUsers, nil
}
