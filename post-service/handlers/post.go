package handlers

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/services"
	"github.com/gorilla/mux"
)

type PostHandler struct {
	Service *services.PostService
}

func (handler *PostHandler) Hello(w http.ResponseWriter, r *http.Request) {
	addrs, _ := net.InterfaceAddrs()
	for i, addr := range addrs {
		fmt.Printf("%d %v\n", i, addr)
	}
}

func (handler *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("creating")
	var post data.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(post)
	err = handler.Service.CreatePost(&post)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	fmt.Println("created")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *PostHandler) Verify(w http.ResponseWriter, r *http.Request) {
	fmt.Println("verifying")
	vars := mux.Vars(r)
	id := vars["picpath"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	exists, err := handler.Service.PostExists(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if exists {
		w.WriteHeader(http.StatusOK)
		fmt.Println("EXISTS")
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
