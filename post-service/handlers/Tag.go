package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/services"
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
