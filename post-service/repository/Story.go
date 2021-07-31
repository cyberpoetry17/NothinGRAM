package repository

import (
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/google/uuid"
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
		if(el.IsActive && el.IsOnlyForCloseFriends==false){
			retList = append(retList, el)
		}
	}
	return retList
}

func (repo *StoryRepo) GetAllUserStories(userId uuid.UUID) ([]data.Story,error){
	var list []data.Story
	err := repo.Database.Where("\"UserId\" = ?",userId).Find(&list).Error
	return list,err
}

func (repo *StoryRepo) GetCloseFrinedStoriesForUser(userId uuid.UUID) ([]data.Story){
	var retList []data.Story
	lis,_ :=repo.GetAllUserStories(userId);
	for _,el:= range lis{
		if(el.IsOnlyForCloseFriends){
			retList = append(retList, el)
		}
	}
	return retList
}

func (repo *StoryRepo) GetUserStoryHighlights(userId uuid.UUID) ([]data.Story){
	var retList []data.Story
	lis,err :=repo.GetAllUserStories(userId)
	if(err!= nil){
		return nil
	}
	for _,el:= range lis{
		if(el.ShowOnStoryHighlights){
			retList = append(retList, el)
		}
	}
	return retList
}

