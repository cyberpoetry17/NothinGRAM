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

func (repo *MutedRepo) GetAll() []data.Muted {
	var muted []data.Muted
	repo.Database.Find(&muted)
	return muted
}

func (repo *MutedRepo) RemoveMuted(mutedUser data.Muted) error {
	mutedUsers := repo.GetAll()
	for _, element := range mutedUsers {
		if element.MutedID.String() == mutedUser.MutedID.String() {
			if element.UserID.String() == mutedUser.UserID.String() {
				repo.Database.Delete(&element)
			}
		}
	}

	// result := repo.Database.Where("userID=? and mutedID=?", mutedUser.UserID, mutedUser.MutedID).Delete(mutedUser)
	// if result.Error != nil {
	// 	return result.Error
	// }

	return nil

}

func (repo MutedRepo) DeleteMutesForUser(userid string) bool {
	mutes := repo.GetAll()
	for _, element := range mutes {
		if element.UserID.String() == userid || element.MutedID.String() == userid {
			repo.Database.Delete(&element)
		}
	}
	return true
}

func (repo MutedRepo) GetAllMutedUsersByID(userID string) []string {
	var result = repo.GetAll()
	var frontList []string
	for _, element := range result {
		if element.MutedID.String() == userID {
			frontList = append(frontList, element.UserID.String())
		}
	}
	return frontList
}

func (repo *MutedRepo) MutedExists(muted uuid.UUID, userid uuid.UUID) bool {
	var count int64
	var exists = false
	repo.Database.Where("mutedID = ? and userID = ?", muted, userid).Find(&data.Muted{}).Count(&count)
	if count > 0 {
		exists = true
	}
	return exists
}

func (repo *MutedRepo) MutedStatusForProfile(muted *data.Muted) bool {
	var result = repo.GetAll()
	for _, element := range result {
		if element.MutedID == muted.MutedID && element.UserID == muted.UserID {
			return true
		}
	}
	return false
}
