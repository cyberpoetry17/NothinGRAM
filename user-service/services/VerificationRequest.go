package services

import (
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/repository"
)

type VerificationRequestService struct {
	Repo *repository.VerificationRequestRepo
}

type UserVerificationRequest struct {
	Name          string             `json:"name"`
	Surname       string             `json:"surname"`
	Username      string             `json:"username"`
	Category      data.Category      `json:"category"`
	RequestStatus data.RequestStatus `json:"status"`
	PicturePath   string             `json:"picture"`
}

func (service *VerificationRequestService) CreateVerificationRequest(verificationRequest *data.VerificationRequest) error {
	service.Repo.CreateVerificationRequest(verificationRequest)
	return nil
}

func (service *VerificationRequestService) VerificationRequestExistsByUsername(username string) bool {
	exists := service.Repo.VerificationRequestExistsByUsername(username)
	return exists
}

func (service *VerificationRequestService) GetAllWaitlistedVerificationRequests() []data.VerificationRequest {
	allVerificationRequests := service.Repo.GetAllVerificationRequests()
	var waitlistedVerificationRequests []data.VerificationRequest
	for _, verificationRequest := range allVerificationRequests {
		if verificationRequest.RequestStatus == 3 {
			waitlistedVerificationRequests = append(waitlistedVerificationRequests, verificationRequest)
		}
	}

	return waitlistedVerificationRequests
}

func (service *VerificationRequestService) UpdateVerificationRequest(verificationRequest *data.VerificationRequest) error {
	errorDelete := service.Repo.Database.Where("username=?", verificationRequest.Username).Delete(verificationRequest).Error
	if errorDelete != nil {
		return errorDelete
	}

	errorSave := service.Repo.Database.Save(verificationRequest).Where("username=?", verificationRequest.Username).Error
	if errorSave != nil {
		return errorSave
	}
	return nil
}
