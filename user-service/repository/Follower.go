package repository

import (
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"gorm.io/gorm"
)

type FollowerRepo struct {
	Database *gorm.DB
}

func (repo *FollowerRepo) FollowUser(follower *data.Follower) error {
	result := repo.Database.Create(follower)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *FollowerRepo) UnfollowUser(follower *data.Follower) error {
	result := repo.Database.Where("idfollower=? and iduser=?", follower.IDFollower, follower.IDUser).Delete(follower)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
func (repo *FollowerRepo) GetAll() []data.Follower{
	var followers []data.Follower
	repo.Database.
		Find(&followers)
	return followers
}
func (repo *FollowerRepo) FollowStatusForProfile(follower *data.Follower) bool {
	var result = repo.GetAll()
	for _,element := range result{
		if element.IDFollower == follower.IDFollower && element.IDUser == follower.IDUser{
			return true
		}
	}
	return false
}
