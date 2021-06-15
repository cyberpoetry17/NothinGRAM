package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/services"
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