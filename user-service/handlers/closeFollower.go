package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/cyberpoetry17/NothinGRAM/UserAPI/DTO"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/services"
	"github.com/dgrijalva/jwt-go"
)

type CloseFollowerHandler struct {
	Service  *services.CloseFollowerService
	UserServ *services.UserService
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

func (handler *CloseFollowerHandler) ModifyCloseFollowers(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	token := r.Header.Get("Authorization")
	splitToken := strings.Split(token, "Bearer ")
	token = splitToken[1]

	tknStr := token
	tokenObj := &data.Token{}
	tkn, err := jwt.ParseWithClaims(tknStr, tokenObj, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusAlreadyReported)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var followerModificationLists DTO.UsernameDTO
	errr := json.NewDecoder(r.Body).Decode(&followerModificationLists)

	if errr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, element := range followerModificationLists.Usernames {
		fmt.Println(element)
	}

	addedUsers := handler.UserServ.SetCloseFollowersToUser(followerModificationLists.Usernames, tokenObj.UserID)

	for _, element := range addedUsers {
		fmt.Println(element)
		fmt.Println("hahaha")
	}
	removedUsers := handler.UserServ.SetCloseFollowersToUser(followerModificationLists.RemovedUsernames, tokenObj.UserID)

	errorDeleted := handler.Service.RemoveManyByID(removedUsers)
	errorAdded := handler.Service.AddMultipleFollowers(addedUsers, tokenObj.UserID)
	if errorAdded != nil || errorDeleted != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

}
