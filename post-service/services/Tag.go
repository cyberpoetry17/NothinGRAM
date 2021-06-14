package services

import (
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/repository"
	"github.com/google/uuid"
)

type TagService struct {
	Repo *repository.TagRepo
}

func (service *TagService) CreateTag(tag *data.Tag) error {
	return  service.Repo.CreateTag(tag)
}

func (service *TagService) TagExists(userId uuid.UUID) bool {
	return service.Repo.TagExists(userId)
}

func (service *TagService) TagExistsByName(tagName string) bool {
	return service.Repo.TagExistsByName(tagName)
}

func (service *TagService) GetTagByName(tagName string) *data.Tag {
	return service.Repo.GetTagByName(tagName)
}

func (service *TagService) EditTag(tag *data.Tag) error {
	return service.Repo.EditTag(tag)
}

func (service *TagService) RemoveTag(tag *data.Tag) error {
	return service.Repo.RemoveTag(tag)
}