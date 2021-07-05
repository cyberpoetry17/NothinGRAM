package repository

import (
	"fmt"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"gorm.io/gorm"
)

type DislikeRepo struct {
	Database *gorm.DB
}
func (repo DislikeRepo) CreateDislike(dislike *data.Dislike) error {
	result := repo.Database.Create(dislike)
	if(result.Error != nil){
		return result.Error
	}
	fmt.Println(result.RowsAffected)
	return nil
}

func (repo DislikeRepo) GetAllDislikesForPost (postId string) []data.Dislike{
	var dislikes []data.Dislike
	var backList []data.Dislike
	repo.Database.Preload("Posts").Find(&dislikes)
	for _,element := range dislikes{
		if element.PostId.String() == postId {
			backList = append(backList,element)
		}
	}
	return backList
}

func (repo DislikeRepo) RemoveDislike (dislike *data.Dislike) error{
	return repo.Database.Delete(dislike).Error
}

func (repo DislikeRepo) CheckIfUserDislikedPost (dislike *data.Dislike) bool{
	var dislikes []data.Dislike
	dislikes = repo.GetAllDislikesForPost(dislike.PostId.String())
	for _,element := range dislikes{
		if element.UserId == dislike.UserId{
			return true
		}
	}
	return false
}