package repository

import (
	"fmt"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MediaRepo struct {
	Database *gorm.DB
}

func (repo *MediaRepo) CreateMedia(media *data.Media) error {
	result := repo.Database.Create(media)
	if(result.Error != nil){
		return result.Error
	}
	fmt.Println(result.RowsAffected)
	return nil
}

func (repo *MediaRepo) RemoveMedia(media *data.Media) error {
	return repo.Database.Delete(media).Error
}

func (repo *MediaRepo) EditMedia(media *data.Media) error {
	return repo.Database.Save(media).Error
}

func (repo *MediaRepo) GetMediaForPost(postID uuid.UUID) (error,[]data.Media) {
	var mediaToReturn []data.Media
	fmt.Println(postID.String())
	err := repo.Database.Where("\"PostId\" = ?",postID.String()).Find(&mediaToReturn).Error
	return err,mediaToReturn
}

func (repo *MediaRepo) GetMediaForStory(storyID uuid.UUID) (error,[]data.Media) {
	var mediaToReturn []data.Media
	fmt.Println(storyID.String())
	err := repo.Database.Where("\"StoryId\" = ?", storyID.String()).Find(&mediaToReturn).Error
	return err,mediaToReturn
}