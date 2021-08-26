package repository

import (
	"fmt"

	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BlockedRepo struct {
	Database *gorm.DB
}

func (repo BlockedRepo) CreateBlocked(blocked *data.Blocked) error {
	result := repo.Database.Create(blocked)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println(result.RowsAffected)
	return nil
}
func (repo BlockedRepo) GetAll() []data.Blocked{
	var blocked []data.Blocked
	repo.Database.Find(&blocked)
	return blocked
}
func (repo BlockedRepo) RemoveBlocked(blocked *data.Blocked) error {
	return repo.Database.Delete(blocked).Error
}

func (repo BlockedRepo) DeleteBlocksForUser(userid string) bool {
	blocked := repo.GetAll()
	for _,element := range blocked{
		if element.UserID.String() == userid || element.BlockedID.String() == userid{
			repo.Database.Delete(&element)
		}
	}
	return true
}

func (repo BlockedRepo) GetAllBlockedUsersByID(userID string) ([]data.Blocked, error) {
	id, err := uuid.Parse(userID)
	if err != nil {
		print(err)
		return nil, err
	}
	var blockedUsers []data.Blocked
	repo.Database.Find(&blockedUsers).Where("userID = ?", id)
	repo.Database.Preload("user2", &blockedUsers)
	return blockedUsers, nil
}
