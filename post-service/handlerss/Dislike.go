package handlerss

import (
	"encoding/json"
	"fmt"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/services"
	"github.com/gorilla/mux"
	"net/http"
)

type DislikeHandler struct {
	Service *services.DislikeService
}

func (handler *DislikeHandler) CreateDislike(w http.ResponseWriter, r *http.Request) {
	fmt.Println("creating dislike")
	var dislike data.Dislike
	err := json.NewDecoder(r.Body).Decode(&dislike)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(dislike)
	err = handler.Service.CreateDislike(&dislike)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *DislikeHandler) DeleteDislike(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deleting dislike")
	var dislike data.Dislike
	err := json.NewDecoder(r.Body).Decode(&dislike)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(dislike)
	err = handler.Service.RemoveDislike(&dislike)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *DislikeHandler) GetAllDislikesForPost (w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	id := vars["postid"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	dislikes := handler.Service.GetAllDislikesForPost(id)

	if len(dislikes)!=0 {
		w.WriteHeader(http.StatusOK)
		for i,dislikes := range dislikes{
			fmt.Println("%d : %s", i,dislikes.IDD)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func (handler *DislikeHandler) CheckIfUserDislikedPost (w http.ResponseWriter,r *http.Request){
	var dislike data.Dislike
	err := json.NewDecoder(r.Body).Decode(&dislike)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	disliked := handler.Service.CheckIfUserDislikedPost(&dislike)

	if disliked != false {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(true)
	} else {
		_ = json.NewEncoder(w).Encode(false)

	}
}