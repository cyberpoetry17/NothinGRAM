package repository

import (
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"gorm.io/gorm"
)

type FollowerRequestRepo struct {
	Database *gorm.DB
}

func (repo *FollowerRequestRepo) CreateFollowRequest(followerRequest *data.FollowerRequest) error {
	result := repo.Database.Create(followerRequest)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *FollowerRequestRepo) GetAllForUser(id string) []data.FollowerRequest{
	var requests []data.FollowerRequest
	var frontList []data.FollowerRequest
	repo.Database.Find(&requests)
	for _,element := range requests{
		if element.IDFollowed.String() == id{
			frontList = append(frontList, element)
		}
	}
	return frontList
}

func (repo *FollowerRequestRepo) RemoveRequest(id string) error {
	var request data.FollowerRequest
	return repo.Database.Where("id = ?",id).Delete(request).Error
}