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
}

func (service *VerificationRequestService) CreateVerificationRequest(verificationRequest *data.VerificationRequest) error {
	service.Repo.CreateVerificationRequest(verificationRequest)
	return nil
}

func (service *VerificationRequestService) VerificationRequestExistsByUsername(username string) bool {
	exists := service.Repo.VerificationRequestExistsByUsername(username)
	return exists
}
