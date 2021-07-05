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
	var backList []data.Like
	repo.Database.Preload("Posts").Find(&likes)
	for _,element := range likes{
		if element.PostId.String() == postId {
			backList = append(backList,element)
		}
	}
	return backList
}

func (repo LikeRepo) RemoveLike (like *data.Like) error{
	return repo.Database.Delete(like).Error
}

func (repo LikeRepo) CheckIfUserLikedPost (like *data.Like) bool{
	var likes []data.Like
	likes = repo.GetAllLikesForPost(like.PostId.String())
	for _,element := range likes{
		if element.UserId == like.UserId{
			return true
		}
	}
	return false
}