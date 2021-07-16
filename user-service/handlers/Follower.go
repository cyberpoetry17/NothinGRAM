package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
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

func (handler *FollowerHandler) FollowStatusForProfile(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w,r)
	var follower data.Follower
	err := json.NewDecoder(r.Body).Decode(&follower)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ret := handler.Service.FollowStatusForProfile(&follower)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
	}
	if ret == true{
		_ = json.NewEncoder(w).Encode(true)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
	}else{
		_ = json.NewEncoder(w).Encode(false)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
	}
}

func (handler *FollowerHandler) FollowedByUser (w http.ResponseWriter,r *http.Request){
	setupResponse(&w,r)
	vars := mux.Vars(r)
	id := vars["userid"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	followed := handler.Service.FollowedByUser(id)


	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(followed)
}
