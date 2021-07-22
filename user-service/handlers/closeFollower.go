package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/services"
)

type CloseFollowerHandler struct {
	Service *services.CloseFollowerService
}

func (handler *CloseFollowerHandler) AddCloseFollower(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	var follower data.CloseFollower
	err := json.NewDecoder(r.Body).Decode(&follower)
	if err != nil {
		fmt.Println("-------------HERE BAD REQUEST-------------")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.AddCloseFollower(&follower)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *CloseFollowerHandler) RemoveCloseFollower(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	var follower data.CloseFollower
	err := json.NewDecoder(r.Body).Decode(&follower)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.RemoveCloseFollower(&follower)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
