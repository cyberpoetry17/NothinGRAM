package repository

import (
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/DTO"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CloseFollowerRepository struct {
	Database *gorm.DB
}

func (repo *CloseFollowerRepository) AddCloseFollower(follower *data.CloseFollower) error {
	result := repo.Database.Create(follower)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *CloseFollowerRepository) RemoveCloseFollower(follower *data.CloseFollower) error {
	result := repo.Database.Where("idfollower=? and iduser=?", follower.IDCloseFollower, follower.IDUser).Delete(follower)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *CloseFollowerRepository) GetAllCloseFollowers() []data.CloseFollower {
	var followers []data.CloseFollower
	repo.Database.
		Find(&followers)
	return followers
}

func (repo *CloseFollowerRepository) GetAllCloseFollowerUser(userid string) []string {
	var result = repo.GetAllCloseFollowers()
	var frontList []string
	for _, element := range result {
		if element.IDCloseFollower.String() == userid {
			frontList = append(frontList, element.IDUser.String())
		}
	}
	return frontList
}

func (repo *CloseFollowerRepository) RemoveManyByID(element DTO.UserDTO, id uuid.UUID) error {
	result := repo.Database.Where("idfollower=? and iduser=?", element.UserId, id).Delete(&data.CloseFollower{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *CloseFollowerRepository) CloseFollowerExists(idclosefollower uuid.UUID, iduser uuid.UUID) int {
	user := data.CloseFollower{}
	err := repo.Database.Where("idfollower=? and iduser=?", idclosefollower, iduser).First(&user).Error
	if err != nil {
		return 0

	}

	return 1
}
