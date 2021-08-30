package repository

import (
	"fmt"

	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"gorm.io/gorm"
)

type VerificationRequestRepo struct {
	Database *gorm.DB
}

func (repo *VerificationRequestRepo) CreateVerificationRequest(verificationRequest *data.VerificationRequest) error {
	result := repo.Database.Create(verificationRequest)
	fmt.Println(result.RowsAffected)
	return nil
}

func (repo *VerificationRequestRepo) VerificationRequestExistsByUsername(username string) bool {
	var count int64
	repo.Database.Where("username=?", username).Find(&data.VerificationRequest{}).Count(&count)
	return count != 0
}

func (repo *VerificationRequestRepo) GetAllVerificationRequests() []data.VerificationRequest {
	var allVerificationRequests []data.VerificationRequest
	repo.Database.
		Preload("Name").
		Preload("Surname").
		Preload("Username").
		Preload("Category").
		Preload("RequestStatus").
		Preload("PicturePath").
		Find(&allVerificationRequests)

	return allVerificationRequests
}
