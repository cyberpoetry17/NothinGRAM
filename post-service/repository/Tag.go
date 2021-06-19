package repository

import (
	"fmt"

	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TagRepo struct {
	Database *gorm.DB
}

func (repo *TagRepo) CreateTag(tag *data.Tag) error {
	result := repo.Database.Create(tag)
	if(result.Error != nil){
		return result.Error
	}
	fmt.Println(result.RowsAffected)
	return nil //sta s ovim nilom
}

//BY ID
func (repo *TagRepo) TagExists(tagId uuid.UUID) bool {
	var count int64
	repo.Database.Where("id = ?", tagId).Find(&data.Tag{}).Count(&count)
	return count != 0
}

func (repo *TagRepo) TagExistsByName(tagName string) bool {
	var count int64
	repo.Database.Where("TagName = ?", tagName).Find(&data.Tag{}).Count(&count)
	return count != 0
}

func (repo *TagRepo) GetTagByName(tagName string) *data.Tag {
	var tag data.Tag
	err := repo.Database.Where("TagName = ?", tagName).First(tag).Error
	if err == nil {
		return nil
	}
	return  &tag
}

func (repo *TagRepo) GetPostByTag(tagId string) []data.Post{
	var posts []data.Post
	var tagFound = repo.GetAllTags()
	for _,element := range tagFound{
		if element.ID.String() ==  tagId{
			for _,el := range element.Posts{
				posts = append(posts, el)
			}
		}
	}
	return posts
}

func (repo *TagRepo) EditTag(tag *data.Tag) error {
	return repo.Database.Save(tag).Error
	//return repo.Database.Model(tag).Update("TagName",tag.TagName).Error
}

func (repo *TagRepo) RemoveTag(tag *data.Tag) error {
	return repo.Database.Delete(tag).Error
}

func (repo *TagRepo) GetAllTags() []data.Tag{
	var tags []data.Tag
	repo.Database.Preload("Posts").Find(&tags)	//check if work !!!
	return tags
}

func (repo *TagRepo) FilterPublicMaterialByTag(tagId string) []data.Post{
	var media []data.Post
	//var frontList []data.Post
	media = repo.GetPostByTag(tagId)
	//for _,element := range media{				prosirenje funkcije za kad se ubaci user
	//	if element.UserID.isPublic(){
	//		frontList = append(frontList, element)
	//	}
	//}
	for _,el := range media{
		fmt.Println(el.ID)
	}
	return media//frontList
}
//func (repo *TagRepo) AddPostToTag(tag *data.Tag, post)error{
//	repo.Database.
//	repo.Database.Model(tag).Update("")
//
//}