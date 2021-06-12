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
	user := &data.User2{}
	//kastujem iz request tela u user strukturu
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := handler.Service.FindOneByEmailAndPassword(user.Email, user.Password)
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

// func JwtVerify(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		var header = r.Header.Get("x-access-token") //Grab the token from the header

// 		header = strings.TrimSpace(header)

// 		if header == "" {
// 			//Token is missing, returns with error code 403 Unauthorized
// 			w.WriteHeader(http.StatusForbidden)
// 			json.NewEncoder(w).Encode(data.Exception{Message: "Missing auth token"})
// 			return
// 		}
// 		tk := &data.Token{}

// 		_, err := jwt.ParseWithClaims(header, tk, func(token *jwt.Token) (interface{}, error) {
// 			return []byte("secret"), nil
// 		})

// 		if err != nil {
// 			w.WriteHeader(http.StatusForbidden)
// 			json.NewEncoder(w).Encode(data.Exception{Message: err.Error()})
// 			return
// 		}

// 		ctx := context.WithValue(r.Context(), "user", tk)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
//}
