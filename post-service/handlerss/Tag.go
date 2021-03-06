package handlerss

import (
	"encoding/json"
	"fmt"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/services"
	"github.com/gorilla/mux"
	"net/http"
)

type TagHandler struct {
	Service *services.TagService
}

func (handler *TagHandler) CreateTag(w http.ResponseWriter, r *http.Request) {
	fmt.Println("creating tag")
	var tag data.Tag
	err := json.NewDecoder(r.Body).Decode(&tag)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(tag)
	err = handler.Service.CreateTag(&tag)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *TagHandler) EditTag(w http.ResponseWriter, r *http.Request) {
	fmt.Println("edit tag")
	var tag data.Tag
	err := json.NewDecoder(r.Body).Decode(&tag)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(tag)
	err = handler.Service.EditTag(&tag)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *TagHandler) DeleteTag(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete tag")
	var tag data.Tag
	err := json.NewDecoder(r.Body).Decode(&tag)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(tag)
	err = handler.Service.RemoveTag(&tag)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *TagHandler) FilterPublicMaterialByTagId(w http.ResponseWriter, r *http.Request) {
	fmt.Println("filtering by tag")
	vars := mux.Vars(r)
	tagid := vars["tagid"]
	if tagid == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(tagid)
	handler.Service.FilterPublicMaterialByTags(tagid)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *TagHandler) GetAllTagNames(w http.ResponseWriter, r *http.Request) {
	tags := handler.Service.GetAllTagsNames()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tags)
}

func (handler *TagHandler) GetAllTags(w http.ResponseWriter, r *http.Request) {
	tags := handler.Service.GetAllTags()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tags)
}
