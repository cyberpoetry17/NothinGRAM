package repository

import (
	"fmt"

	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepo struct {
	Database *gorm.DB
}

func (repo *UserRepo) CreateUser(user *data.User2) error {
	result := repo.Database.Create(user)
	fmt.Println(result.RowsAffected)
	return nil
}
func (repo *UserRepo) GetById(id uuid.UUID) (*data.User2, error) {
	user := &data.User2{}

	err := repo.Database.Where("id = ?", id).Preload("blocked").First(user).Error

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepo) UserExists(userId uuid.UUID) bool {
	var count int64
	repo.Database.Where("id = ?", userId).Find(&data.User2{}).Count(&count)
	return count != 0
}

func (repo *UserRepo) UserExistsByEmail(email string) bool {
	var count int64
	repo.Database.Where("email = ?", email).Find(&data.User2{}).Count(&count)
	return count != 0
}

func (repo *UserRepo) UserExistsByUsername(username string) bool {
	var count int64
	repo.Database.Where("username = ?", username).Find(&data.User2{}).Count(&count)
	return count != 0
}

// func (repo *UserRepo) AddTagToPost(tag data.Tag,userID uuid.UUID) error{
// 	for _, element := range repo.GetAllBlockedUsers(){
// 		if(element.ID == userID){
// 			element.Tags = append(element.Tags, tag)
// 			//repo.Database.Model(&data.Post{}).Association("Tags").Append(tag)
// 			////return nil
// 			//err := repo.Database.Save(&element).Error	//ovo radi ali kreira novi tag

// 			//err:=repo.Database.Session(&gorm.Session{FullSaveAssociations: true}).Save(element).Error
// 			err := repo.Database.Raw("INSERT INTO posts_tags (tag_id,post_id) VALUES (?,?)",tag.ID.String(),element.ID.String()).Error

// 			return err
// 		}
// 	}
// 	return nil
// }
