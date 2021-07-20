package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/services"
	"github.com/gorilla/mux"
	"net/http"
)

type FollowerRequestHandler struct {
	Service *services.FollowerRequestService
}

func (handler *FollowerRequestHandler) CreateFollowRequest(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w,r)
	var followerRequest data.FollowerRequest
	err := json.NewDecoder(r.Body).Decode(&followerRequest)
	if err != nil {
		fmt.Println("-------------HERE BAD REQUEST-------------")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.CreateFollowRequest(&followerRequest)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *FollowerRequestHandler) GetAllRequests(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w,r)
	fmt.Println("Getting all requests..")
	vars := mux.Vars(r)
	id := vars["userid"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(handler.Service.GetAllRequests(id))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}

func (handler *FollowerRequestHandler) DeleteRequest(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w,r)
	fmt.Println("deleting request")
	vars := mux.Vars(r)
	id := vars["requestid"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := handler.Service.RemoveRequest(id)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}