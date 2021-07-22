package repository

import (
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
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
