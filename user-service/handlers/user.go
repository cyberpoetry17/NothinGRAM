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

type UserHandler struct {
	Service *services.UserService
}

func (handler *UserHandler) Hello(w http.ResponseWriter, r *http.Request) {
	addrs, _ := net.InterfaceAddrs()
	for i, addr := range addrs {
		fmt.Printf("%d %v\n", i, addr)
	}
}

//login user
func (handler *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	//user := &data.User2{}
	var userRequest services.LoginRequest
	//kastujem iz request tela u user strukturu
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
		//TODO log
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(user)
	existsByUsername := handler.Service.Repo.UserExistsByUsername(user.Username)
	existsByEmail := handler.Service.Repo.UserExistsByEmail(user.Email)

	//Fault User Already Exits
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
		fmt.Println("aaaaaaaaaaa!")
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	//kastujem iz request tela u user strukturu

	// fmt.Println(user.ID)
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

// func (u *data.User2) updateMe(w http.ResponseWriter, r *http.Request) error {
// 	me := domain.UserMustFromContext(r.Context())
// 	req := new(engine.UpdateUserRequest)
// 	if err := decodeReq(r, req); err != nil {
// 		return err
// 	}

// 	req.ID = me.ID

// 	if err := u.Update(req); err != nil {
// 		if err == engine.ErrEmailExists {
// 			return newWebErr(emailExistsErrCode, http.StatusUnprocessableEntity, err)
// 		}
// 		return err
// 	}

// 	gores.NoContent(w)
// 	return nil
// }
