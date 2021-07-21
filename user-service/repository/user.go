package repository

import (
	"fmt"

	"github.com/cyberpoetry17/NothinGRAM/UserAPI/DTO"

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

func (repo *UserRepo) GetAll() []data.User2 {
	var users []data.User2
	repo.Database.
		Preload("Followers").
		Preload("Following").
		Preload("MutedUsers").
		Preload("BlockedUsers").
		Find(&users)
	return users
}

func (repo *UserRepo) GetById(id uuid.UUID) (*data.User2, error) {
	user := &data.User2{}

	err := repo.Database.
		Preload("Followers").
		Preload("Following").
		Preload("MutedUsers").
		Preload("BlockedUsers").
		Where("id = ?", id).First(user).Error

	if err != nil {
		return nil, err
	}
	return user, nil
}

// func (repo *UserRepo) GetById(id uuid.UUID) (*data.User2,error) {
// 	var users []data.User2
// 	var backUser data.User2
// 	users = repo.GetAll()
// 	for _,element := range users{
// 		if element.ID == id {
// 			backUser = element
// 			return &backUser,nil
// 		}
// 	}
// 	return &backUser,nil
// }

func (repo *UserRepo) GetUserByUsernameForProfile(id string) *data.User2 {
	var users []data.User2
	var backUser data.User2
	users = repo.GetAll()
	for _, element := range users {
		if element.Username == id {
			backUser = element
			return &backUser
		}
	}
	return &backUser
}

func (repo *UserRepo) GetUserIdByUsernameForProfile(id string) DTO.UserUsernameAndPrivateDTO {
	var users []data.User2
	var backUser DTO.UserUsernameAndPrivateDTO
	users = repo.GetAll()
	for _, element := range users {
		if element.Username == id {
			backUser.UserId = element.ID
			backUser.Private = element.Private
			return backUser
		}
	}
	return backUser
}

func (repo *UserRepo) GetUsernameById(id uuid.UUID) string {
	var users []data.User2
	var backUser string
	users = repo.GetAll()
	for _, element := range users {
		if element.ID == id {
			backUser = element.Username
			return backUser
		}
	}
	return backUser
}

func (repo *UserRepo) GetByEmail(email string) (*data.User2, error) {
	user := &data.User2{}

	err := repo.Database.Where("email = ?", email).Preload("blocked").First(user).Error

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
