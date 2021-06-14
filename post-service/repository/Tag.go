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
func (repo *TagRepo) TagExists(userId uuid.UUID) bool {
	var count int64
	repo.Database.Where("id = ?", userId).Find(&data.Tag{}).Count(&count)
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
