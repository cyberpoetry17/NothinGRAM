package services

import (
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/repository"
	"github.com/google/uuid"
)

type PostService struct {
	Repo *repository.PostRepo
}

//verovatno treba da vrati neku vrednost
func (service *PostService) CreatePost(post *data.Post) error {
	return service.Repo.CreatePost(post)
}

func (service *PostService) PostExists(desc string) (bool, error) {
	id := desc
	exists := service.Repo.PostExists(id)
	return exists, nil
}

func (service *PostService) AddTagToPost(tag data.Tag,postId uuid.UUID) error{
	return service.Repo.AddTagToPost(tag,postId)
}

func (service *PostService) AddLocationToPost(location data.Location,postId uuid.UUID) error{
	return service.Repo.AddLocationToPost(location,postId)
}
func (service *PostService) GetAllPosts() []data.Post{
	return service.Repo.GetAll()
}

func (service *PostService) GetPostsByUserID(id string) []data.Post{
	return service.Repo.GetPostsByUserID(id)
}

func (service *PostService) GetUsernameByPostUserID(userid string) string{
	return service.Repo.GetUsernameByPostUserID(userid)
}