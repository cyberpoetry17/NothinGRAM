package repository

import (
	"fmt"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"gorm.io/gorm"
)

type LikeRepo struct {
	Database *gorm.DB
}

func (repo LikeRepo) CreateLike(like *data.Like) error {
	result := repo.Database.Create(like)
	if(result.Error != nil){
		return result.Error
	}
	fmt.Println(result.RowsAffected)
	return nil
}

func (repo LikeRepo) GetAllLikesForPost (postId string) []data.Like{
	var likes []data.Like
	repo.Database.Find(&likes).Where("PostId = ?",postId)
	repo.Database.Preload("Posts",&likes)
	return likes
}

func (repo LikeRepo) RemoveLike (like *data.Like) error{
	return repo.Database.Delete(like).Error
}