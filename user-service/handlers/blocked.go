package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/services"
)

type BlockedHandler struct {
	Service *services.BlockedService
}

func (handler *BlockedHandler) BlockUser(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	fmt.Println("blocking user..")
	var blockedUser data.Blocked
	err := json.NewDecoder(r.Body).Decode(&blockedUser)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(blockedUser.BlockedID)
	fmt.Print("useeer id blocked")
	err = handler.Service.CreateBlockedUser(&blockedUser)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *BlockedHandler) UnblockUser(w http.ResponseWriter, r *http.Request) { //DELETING BLOCKED USER STRUCT
	fmt.Println("unblocking user...")
	var blockedUser data.Blocked
	err := json.NewDecoder(r.Body).Decode(&blockedUser)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(blockedUser)
	err = handler.Service.RemoveBlockedUser(&blockedUser)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

// func (handler *BlockedHandler) GetAllBlockedUsers(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["userID"]
// 	if id == "" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	blockedUsers, error := handler.Service.GetAllBlockedUsers(id)
// 	if error != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	if len(blockedUsers) != 0 {
// 		w.WriteHeader(http.StatusOK)
// 		for i, blockedUsers := range blockedUsers {
// 			fmt.Printf("%d : %s", i, blockedUsers.BlockedID)
// 		}
// 	} else {
// 		w.WriteHeader(http.StatusNotFound)
// 	}
// }

func (handler *BlockedHandler) BlockStatusForProfile(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	var blocked data.Blocked
	err := json.NewDecoder(r.Body).Decode(&blocked)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ret := handler.Service.BlockedStatusForProfile(&blocked)
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
