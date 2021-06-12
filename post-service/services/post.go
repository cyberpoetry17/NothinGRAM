package services

import (
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/repository"
)

type PostService struct {
	Repo *repository.PostRepo
}

//verovatno treba da vrati neku vrednost
func (service *PostService) CreatePost(post *data.Post) error {
	service.Repo.CreatePost(post)
	return nil
}

func (service *PostService) PostExists(desc string) (bool, error) {
	id := desc
	exists := service.Repo.PostExists(id)
	return exists, nil
}
