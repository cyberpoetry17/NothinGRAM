package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/services"
)

type VerificationRequestHandler struct {
	Service     *services.VerificationRequestService
	ServiceUser *services.UserService
}

func CreateVerificationRequestFromDTO(dto services.UserVerificationRequest) *data.VerificationRequest {
	var verificationRequest data.VerificationRequest
	verificationRequest.Name = dto.Name
	verificationRequest.Surname = dto.Surname
	verificationRequest.Username = dto.Username
	verificationRequest.Category = dto.Category
	verificationRequest.RequestStatus = 3
	verificationRequest.PicturePath = dto.PicturePath

	return &verificationRequest
}

func (handler *VerificationRequestHandler) CreateVerificationRequest(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if (*r).Method == "OPTIONS" {
		return
	}

	var userVerificationRequest services.UserVerificationRequest
	err := json.NewDecoder(r.Body).Decode(&userVerificationRequest)

	if err != nil {
		println(err)
		w.WriteHeader(http.StatusNoContent)
		return
	}

	verificationRequest := CreateVerificationRequestFromDTO(userVerificationRequest)
	verificationRequestExists := handler.Service.Repo.VerificationRequestExistsByUsername(verificationRequest.Username)

	if verificationRequestExists {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.Service.CreateVerificationRequest(verificationRequest)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	fmt.Println("Verification request created.")
	w.WriteHeader(http.StatusCreated)
}

func (handler *VerificationRequestHandler) GetAllWaitlistedVerificationRequests(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if (*r).Method == "OPTIONS" {
		return
	}

	waitlistedVerificationRequests := handler.Service.GetAllWaitlistedVerificationRequests()
	if waitlistedVerificationRequests == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(waitlistedVerificationRequests)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (handler *VerificationRequestHandler) AcceptUserVerificationRequest(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if (*r).Method == "OPTIONS" {
		return
	}

	var userVerificationRequest services.UserVerificationRequest
	err := json.NewDecoder(r.Body).Decode(&userVerificationRequest)

	if err != nil {
		println(err)
		w.WriteHeader(http.StatusNoContent)
		return
	}

	verificationRequest := CreateVerificationRequestFromDTO(userVerificationRequest)
	verificationRequest.RequestStatus = 1
	user := handler.ServiceUser.GetUserByUsernameForProfile(verificationRequest.Username)
	user.Verified = true
	errorUser := handler.ServiceUser.Repo.Database.Save(&user).Error
	if errorUser != nil {
		fmt.Println(errorUser)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	err = handler.Service.UpdateVerificationRequest(verificationRequest)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *VerificationRequestHandler) DeclineUserVerificationRequest(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if (*r).Method == "OPTIONS" {
		return
	}

	var userVerificationRequest services.UserVerificationRequest
	err := json.NewDecoder(r.Body).Decode(&userVerificationRequest)

	if err != nil {
		println(err)
		w.WriteHeader(http.StatusNoContent)
		return
	}

	verificationRequest := CreateVerificationRequestFromDTO(userVerificationRequest)
	verificationRequest.RequestStatus = 2

	err = handler.Service.UpdateVerificationRequest(verificationRequest)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusOK)
}
