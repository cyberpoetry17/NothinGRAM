package repository

import (
	"fmt"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommentRepo struct {
	Database *gorm.DB
}

func (repo *CommentRepo) CreateComment(comment *data.Comment) error {
	result := repo.Database.Create(comment)
	if(result.Error != nil){
		return result.Error
	}
	fmt.Println(result.RowsAffected)
	return nil //sta s ovim nilom.
}

func (repo *CommentRepo) CommentExists(commentId uuid.UUID) bool {
	var count int64
	repo.Database.Where("id = ?", commentId).Find(&data.Comment{}).Count(&count)
	return count != 0
}

func (repo *CommentRepo) EditComment(comment *data.Comment) error {
	return repo.Database.Save(comment).Error
}

func (repo *CommentRepo) RemoveComment(comment *data.Comment) error {
	return repo.Database.Delete(comment).Error
}

func (repo *CommentRepo) GetAllComments() []data.Comment{
	var comments []data.Comment
	repo.Database.Find(&comments)
	return comments
}

func (repo *CommentRepo) GetAllByPostId(postid string) ([]data.Comment,error){
	var comments []data.Comment
	var frontList []data.Comment
	result := repo.Database.Find(&comments)
	for _,element := range comments{
		if element.PostId.String() == postid{
			frontList = append(frontList,element)
		}
	}
	return frontList,result.Error
}

func (repo *CommentRepo) RemoveAllCommentsForPost (id string) bool{
	comments, _ := repo.GetAllByPostId(id)
	for _,element := range comments{
		repo.Database.Delete(&element)
	}
	return true
}