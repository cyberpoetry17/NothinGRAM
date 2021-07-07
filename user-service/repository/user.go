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

	err := repo.Database.Where("id = ?", id).First(user).Error

	if err != nil {
		return nil, err
	}
	return user, nil
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
