package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/services"
)

type MutedHandler struct {
	Service *services.MutedService
}

func (handler *MutedHandler) CreateMutedUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("muting user..")
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
	fmt.Println("removing muted user...")
	var mutedUser data.Muted
	err := json.NewDecoder(r.Body).Decode(&mutedUser)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(mutedUser)
	err = handler.Service.RemoveMutedUser(&mutedUser)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
