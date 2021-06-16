package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/services"
	"net/http"
)

type MediaHandler struct {
	Service *services.MediaService
}

func (handler *MediaHandler) CreateMedia(w http.ResponseWriter, r *http.Request) {
	fmt.Println("creating tag")
	var media data.Media
	err := json.NewDecoder(r.Body).Decode(&media)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(media)
	err = handler.Service.CreateMedia(&media)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *MediaHandler) EditMedia(w http.ResponseWriter, r *http.Request) {
	fmt.Println("creating tag")
	var media data.Media
	err := json.NewDecoder(r.Body).Decode(&media)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(media)
	err = handler.Service.EditMedia(&media)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *MediaHandler) RemoveMedia(w http.ResponseWriter, r *http.Request) {
	fmt.Println("creating tag")
	var media data.Media
	err := json.NewDecoder(r.Body).Decode(&media)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(media)
	err = handler.Service.RemoveMedia(&media)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}