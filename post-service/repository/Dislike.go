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
	repo.Database.Find(&dislikes).Where("PostId = ?",postId)
	repo.Database.Preload("Posts",&dislikes)
	return dislikes
}

func (repo DislikeRepo) RemoveDislike (dislike *data.Dislike) error{
	return repo.Database.Delete(dislike).Error
}