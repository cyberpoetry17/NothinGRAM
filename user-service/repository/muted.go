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

func (repo MutedRepo) GetAll() []data.Muted{
	var muted []data.Muted
	repo.Database.Find(&muted)
	return muted
}

func (repo MutedRepo) RemoveMuted(muted *data.Muted) error {
	return repo.Database.Delete(muted).Error
}

func (repo MutedRepo) DeleteMutesForUser(userid string) bool {
	mutes := repo.GetAll()
	for _,element := range mutes{
		if element.UserID.String() == userid || element.MutedID.String() == userid{
			repo.Database.Delete(&element)
		}
	}
	return true
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
