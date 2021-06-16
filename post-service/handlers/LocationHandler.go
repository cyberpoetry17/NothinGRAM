package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/services"
	"github.com/gorilla/mux"
	"net/http"
)

type LocationHandler struct {
	Service *services.LocationService
}

func (handler *LocationHandler) CreateLocation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("creating location")
	var location data.Location
	err := json.NewDecoder(r.Body).Decode(&location)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(location)
	err = handler.Service.CreateLocation(&location)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *LocationHandler) DeleteLocation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deleting location")
	var location data.Location
	err := json.NewDecoder(r.Body).Decode(&location)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(location)
	err = handler.Service.RemoveLocation(&location)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *LocationHandler) GetLocationForPost (w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	id := vars["postid"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	location := handler.Service.GetLocationForPost(id)

	if location != nil {
		w.WriteHeader(http.StatusOK)
		fmt.Println("Location id : %s",location.IDLoc)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}