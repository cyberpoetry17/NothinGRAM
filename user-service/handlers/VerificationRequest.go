package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/services"
)

type VerificationRequestHandler struct {
	Service *services.VerificationRequestService
}

/*
func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}*/

func CreateVerificationRequestFromDTO(dto services.UserVerificationRequest) *data.VerificationRequest {
	var verificationRequest data.VerificationRequest
	verificationRequest.Name = dto.Name
	verificationRequest.Surname = dto.Surname
	verificationRequest.Username = dto.Username
	verificationRequest.Category = dto.Category
	verificationRequest.RequestStatus = 3

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
