package repository

import (
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"gorm.io/gorm"
	"time"
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
	for _, el:= range(storyList){
		var endTimeForStory time.Time
		endTimeForStory = el.Time
		endTimeForStory = endTimeForStory.Add(24 * time.Hour)
		if(endTimeForStory.Before(time.Now())){
			el.IsActive = false
			repo.EditStory(&el)
		}
	}
	return storyList
}

func (repo *StoryRepo) GetAllActive() []data.Story{
	var retList []data.Story
	for _,el := range repo.GetAll(){
		if(el.IsActive){
			retList = append(retList, el)
		}

	}
	return retList
}
