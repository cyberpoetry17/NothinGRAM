package services

import (
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/repository"
	"github.com/google/uuid"
)

type CommentService struct {
	Repo *repository.CommentRepo
}

func (repo *CommentService) CreateComment(comment *data.Comment) error {
	return repo.Repo.CreateComment(comment)
}

func (repo *CommentService) CommentExists(commentId uuid.UUID) bool {
	return repo.Repo.CommentExists(commentId)
}

func (repo *CommentService) EditComment(comment *data.Comment) error {
	return repo.Repo.EditComment(comment)
}

func (repo *CommentService) RemoveComment(comment *data.Comment) error {
	return repo.Repo.RemoveComment(comment)
}

func (repo *CommentService) GetAllComments() []data.Comment{
	return repo.Repo.GetAllComments()
}

func (repo *CommentService) GetAllByPostId(postid string) ([]data.Comment,error){
	comments,err := repo.Repo.GetAllByPostId(postid)
	if err != nil{
		return nil,err
	}
	return comments,err

}