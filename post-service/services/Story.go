package services

import (
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/DTO"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/repository"
	"github.com/google/uuid"
	"strings"
	"time"
)

type StoryService struct {
	StoryRepo *repository.StoryRepo
	MediaRepo *repository.MediaRepo
	PostRepo *repository.PostRepo
}

func (service *StoryService) CreateStory(storyDTO *DTO.StoryMediaDTO) error {
	var story data.Story
	story.Type = data.MediaT
	story.Time = time.Now()
	story.UserId = storyDTO.UserId
	story.IsActive = true
	story.IsOnlyForCloseFriends = storyDTO.IsOnlyForCloseFriends
	err:= service.StoryRepo.CreateStory(&story)

	if(err!= nil){
		return err
	}
	var media data.Media
	media.Link = storyDTO.MediaPath
	media.Type =data.Picture
	media.StoryId = &story.IdStory
	setIfItIsVideo(media)
	err =service.MediaRepo.CreateMedia(&media)
	if(err!=nil){
		return err
	}
	story.MediaID = media.ID
	err = service.StoryRepo.EditStory(&story)

	return err
}

func setIfItIsVideo(media data.Media) {
	for _, e := range extensions {
		split := strings.Split(media.Link, "?")
		if len(split) == 0 {
			continue
		}
		if strings.HasSuffix(split[0], e) {
			media.Type = data.Video
			break
		}
	}
}

func (service *StoryService) GetAllActiveStories() []data.Story {
	return service.StoryRepo.GetAllActive()
}

func (service *StoryService) GetAllUserStories(userId uuid.UUID) ([]data.Story,error){
	return service.StoryRepo.GetAllUserStories(userId)
}

func (service *StoryService) GetCloseFrinedStoriesForUser(userId uuid.UUID) ([]data.Story){
	return  service.StoryRepo.GetCloseFrinedStoriesForUser(userId)
}

func (service *StoryService) GetAllStoryHighlights(userId uuid.UUID) []data.Story{
	return service.GetAllStoryHighlights(userId)
}

func (service *StoryService) AddToStoryHighlights(storyId uuid.UUID) error{
	for _,el:= range service.StoryRepo.GetAll(){
		if(el.IdStory == storyId){
			el.ShowOnStoryHighlights=true;
			err:= service.StoryRepo.EditStory(&el);
			return err
		}
	}
	return nil
}

func (service *StoryService) RemoveFromStoryHighlights(storyId uuid.UUID) error{
	for _,el:= range service.StoryRepo.GetAll(){
		if(el.IdStory == storyId){
			el.ShowOnStoryHighlights=false;
			err:= service.StoryRepo.EditStory(&el);
			return err
		}
	}
	return nil
}