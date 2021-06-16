package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/services"
	"github.com/gorilla/mux"
	"net/http"
)

type LikeHandler struct {
	Service *services.LikeService
}

func (handler *LikeHandler) CreateLike(w http.ResponseWriter, r *http.Request) {
	fmt.Println("creating like")
	var like data.Like
	err := json.NewDecoder(r.Body).Decode(&like)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(like)
	err = handler.Service.CreateLike(&like)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *LikeHandler) DeleteLike(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deleting like")
	var like data.Like
	err := json.NewDecoder(r.Body).Decode(&like)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(like)
	err = handler.Service.RemoveLike(&like)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *LikeHandler) GetAllLikesForPost (w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	id := vars["postid"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	likes := handler.Service.GetAllLikesForPost(id)

	if len(likes)!=0 {
		w.WriteHeader(http.StatusOK)
		for i,likes := range likes{
			fmt.Println("%d : %s", i,likes.IDL)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}