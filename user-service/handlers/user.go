package handlers

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/services"
	"github.com/google/uuid"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	Service *services.UserService
}

func (handler *UserHandler) Hello(w http.ResponseWriter, r *http.Request) {
	addrs, _ := net.InterfaceAddrs()
	for i, addr := range addrs {
		fmt.Printf("%d %v\n", i, addr)
	}
}

func (handler *UserHandler) GetById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getById")
	vars := mux.Vars(r)
	id := vars["userId"]

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	idUser, errorParsing := uuid.Parse(id)
	if errorParsing != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp, errorUserGetting := handler.Service.GetUserById(idUser)
	if errorUserGetting != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(resp)
	w.WriteHeader(http.StatusOK)
}

//login user
func (handler *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var userRequest services.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := handler.Service.LoginUser(&userRequest)
	json.NewEncoder(w).Encode(resp)
}

//register user
func (handler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("creating")

	var user data.User2
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(user)
	existsByUsername := handler.Service.Repo.UserExistsByUsername(user.Username)
	existsByEmail := handler.Service.Repo.UserExistsByEmail(user.Email)

	if existsByEmail || existsByUsername {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.CreateUser(&user)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	fmt.Println("Created.")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *UserHandler) Verify(w http.ResponseWriter, r *http.Request) {
	fmt.Println("verifying")
	vars := mux.Vars(r)
	id := vars["userId"]

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	exists, err := handler.Service.UserExists(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if exists {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func (handler *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updating")
	var updateUserRequest services.UpdateUserRequest

	err := json.NewDecoder(r.Body).Decode(&updateUserRequest)
	fmt.Println(updateUserRequest.ID)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	fmt.Print(err)

	err = handler.Service.UpdateEditUser(&updateUserRequest) //ovde saljem update User request
	if err != nil {
		fmt.Println(err)

		w.WriteHeader(http.StatusExpectationFailed)
	}
	fmt.Println("Updated.")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
