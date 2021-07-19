package repository

import (
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"gorm.io/gorm"
)

type StoryRepo struct {
	Database *gorm.DB
}

func (repo *StoryRepo) CreateStory(storyDto *data.Story) (error) {
	return repo.Database.Create(storyDto).Error
}

func (repo *StoryRepo) EditStory(story *data.Story) error  {
	return repo.Database.Save(story).Error
}

func (repo *StoryRepo) RemoveTag(story *data.Story) error {
	return repo.Database.Delete(story).Error
}

func (repo *StoryRepo) GetAll() []data.Story{
	var storyList []data.Story
	repo.Database.
		Preload("Post").
		Preload("Media").
		Find(&storyList)
	return storyList
}
