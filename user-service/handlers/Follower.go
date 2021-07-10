package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/services"
)

type FollowerHandler struct {
	Service *services.FollowerService
}

func (handler *FollowerHandler) FollowUser(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w,r)
	var follower data.Follower
	err := json.NewDecoder(r.Body).Decode(&follower)
	if err != nil {
		fmt.Println("-------------HERE BAD REQUEST-------------")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.FollowUser(&follower)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *FollowerHandler) UnfollowUser(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w,r)
	var follower data.Follower
	err := json.NewDecoder(r.Body).Decode(&follower)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.UnfollowUser(&follower)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
