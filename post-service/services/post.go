package services

import (
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/repository"
	"github.com/google/uuid"
)

type PostService struct {
	PostRepo *repository.PostRepo
	TagRepo  *repository.TagRepo
}

//verovatno treba da vrati neku vrednost
func (service *PostService) CreatePost(post *data.Post) error {
	return service.PostRepo.CreatePost(post)
}

func (service *PostService) PostExists(desc string) (bool, error) {
	id := desc
	exists := service.PostRepo.PostExists(id)
	return exists, nil
}

func (service *PostService) AddTagToPost(tag data.Tag,postId uuid.UUID) error{
	var tagToAdd data.Tag
	service.TagRepo.Database.Where("id = ?",tag.ID).Find(&tagToAdd)
	return service.PostRepo.AddTagToPost(&tagToAdd,postId)
}

func (service *PostService) AddLocationToPost(location data.Location,postId uuid.UUID) error{
	return service.PostRepo.AddLocationToPost(location,postId)
}
func (service *PostService) GetAllPosts() []data.Post{
	return service.PostRepo.GetAll()
}

func (service *PostService) GetNonPrivatePosts() []data.Post{
	return service.PostRepo.GetNonPrivatePosts()
}

func (service *PostService) GetPostsByUserID(id string) []data.Post{
	return service.PostRepo.GetPostsByUserID(id)
}

func (service *PostService) GetUsernameByPostUserID(userid string) string{
	return service.PostRepo.GetUsernameByPostUserID(userid)
}