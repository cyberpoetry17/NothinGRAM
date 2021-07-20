package handlerss

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

	err,id := handler.Service.CreateLocation(&location)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(id)
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


	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(location)
}

func (handler *LocationHandler) FilterPublicMaterialByLocationId(w http.ResponseWriter, r *http.Request) {
	fmt.Println("filtering by location")
	vars := mux.Vars(r)
	locationid := vars["locationid"]
	if locationid == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(locationid)
	handler.Service.FilterPublicMaterialByLocations(locationid)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}