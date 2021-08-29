package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/services"
	"github.com/google/uuid"
)

type MutedHandler struct {
	Service *services.MutedService
}

type MutedRequest struct {
	Token     string
	MutedUser uuid.UUID
}

func (handler *MutedHandler) CreateMutedUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("muting user..")
	setupResponse(&w, r)
	var mutedUser data.Muted
	err := json.NewDecoder(r.Body).Decode(&mutedUser)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(mutedUser)
	err = handler.Service.CreateMutedUser(&mutedUser)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *MutedHandler) RemoveMutedUser(w http.ResponseWriter, r *http.Request) { //DELETING BLOCKED USER STRUCT
	setupResponse(&w, r)
	fmt.Println("removing muted user...")
	var mutedUser data.Muted
	err := json.NewDecoder(r.Body).Decode(&mutedUser)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(mutedUser)
	err = handler.Service.RemoveMutedUser(mutedUser)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *MutedHandler) MutedStatusForProfile(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	var muted data.Muted
	err := json.NewDecoder(r.Body).Decode(&muted)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ret := handler.Service.MuteStatusForProfile(&muted)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
	}
	if ret == true {
		_ = json.NewEncoder(w).Encode(true)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
	} else {
		_ = json.NewEncoder(w).Encode(false)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
	}
}

// func (handler *MutedHandler) GetAllMutedUsers(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["userID"]
// 	if id == "" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	mutedUsers, error := handler.Service.GetAllMutedUsers(id)
// 	if error != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	if len(mutedUsers) != 0 {
// 		w.WriteHeader(http.StatusOK)
// 		for i, mutedUsers := range mutedUsers {
// 			fmt.Printf("%d : %s", i, mutedUsers.MutedID)
// 		}
// 	} else {
// 		w.WriteHeader(http.StatusNotFound)
// 	}
// }
