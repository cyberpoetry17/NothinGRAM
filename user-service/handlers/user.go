package handlers

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"strings"
	"time"

	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	Service *services.UserService
}

type authorizationID struct {
	Token string `json:"token"`
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func (handler *UserHandler) Hello(w http.ResponseWriter, r *http.Request) {
	addrs, _ := net.InterfaceAddrs()
	for i, addr := range addrs {
		fmt.Printf("%d %v\n", i, addr)
	}
}

// func (handler *UserHandler) AuthorizationToken(w http.ResponseWriter, r *http.Request) {
// 	c, err := r.Cookie("token")
// 	if err != nil {
// 		if err == http.ErrNoCookie {
// 			// If the cookie is not set, return an unauthorized status
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}
// 		// For any other type of error, return a bad request status
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	tknStr := c.Value
// 	token := &data.Token{}
// 	tkn, err := jwt.ParseWithClaims(tknStr, token, func(token *jwt.Token) (interface{}, error) {
// 		return []byte("secret"), nil
// 	})
// 	if err != nil {
// 		if err == jwt.ErrSignatureInvalid {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	if !tkn.Valid {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}

// 	w.Write([]byte(fmt.Sprintf("Welcome %s!", token.Username)))
// }

func (handler *UserHandler) AuthorizationToken(w http.ResponseWriter, r *http.Request) {

	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	var tokenStruct authorizationID
	err := json.NewDecoder(r.Body).Decode(&tokenStruct) //ovde se nalazi token sa informacijama
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	tknStr := tokenStruct.Token
	tokenObj := &data.Token{}
	tkn, err := jwt.ParseWithClaims(tknStr, tokenObj, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if tokenObj.ExpiresAt < time.Now().UTC().Local().Unix() {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write([]byte(fmt.Sprintf("Welcome %s!", tokenObj.Username)))
}

func (handler *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {

	id, err := uuid.Parse("kjkszpj")
	if err != nil {
		print(err)

	}
	tk := &data.Token{
		UserID:   id,
		Username: "anonymous",
		Email:    "anonymous",
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(-5 * time.Hour).Local().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tokenString)
}

func (handler *UserHandler) GetById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getById")
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	token := r.Header.Get("Authorization")
	splitToken := strings.Split(token, "Bearer ")
	token = splitToken[1]

	tknStr := token
	tokenObj := &data.Token{}
	tkn, err := jwt.ParseWithClaims(tknStr, tokenObj, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	resp, errorUserGetting := handler.Service.GetUserById(tokenObj.UserID)

	if errorUserGetting != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}
	println(resp.Email)
	json.NewEncoder(w).Encode(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (handler *UserHandler) GetUserByUsernameForProfile(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	fmt.Println("getById")
	vars := mux.Vars(r)
	id := vars["username"]

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	idUser, errorParsing := uuid.Parse(id)
	if errorParsing != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp := handler.Service.GetUserByUsernameForProfile(idUser)
	json.NewEncoder(w).Encode(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (handler *UserHandler) GetUsernameById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getting username By Id")
	vars := mux.Vars(r)
	id := vars["usernamebyid"]

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	idUser, errorParsing := uuid.Parse(id)
	if errorParsing != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp := handler.Service.GetUsernameById(idUser)

	json.NewEncoder(w).Encode(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

//login user
func (handler *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {

	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	var userRequest services.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := handler.Service.LoginUser(&userRequest)
	messageStatus := resp["message"].(string)

	if messageStatus == "Email address not found" || messageStatus == "Invalid login credentials. Please try again" {
		w.WriteHeader(http.StatusUnauthorized)
		var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials. Please try again"}
		json.NewEncoder(w).Encode(resp)
		return

	}
	tokenString := resp["token"].(string)
	//expirationTime := resp["expirationDate"].(time.Time)
	println("token string: \n")
	println(tokenString)

	fmt.Println("aaaaaaaaa")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func parseString(info string) time.Time {
	s := strings.Split(info, ":")
	println(s)
	firstPart := s[0]
	println(firstPart)
	dateString := ""
	//runes := []rune(firstPart)
	for i := 0; i < len(firstPart)-3; i++ {
		dateString += string(firstPart[i])

	}
	println(dateString)
	t, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		println("Time parsing not supported!")
	}
	return t
}

func createUserFromDTO(dto services.RegisterRequest, date time.Time) *data.User2 {
	var user data.User2
	user.DateOfBirth = date
	user.Email = dto.Email
	user.Name = dto.Name
	user.Gender = dto.Gender
	user.Password = dto.Password
	user.PhoneNumber = dto.PhoneNumber
	user.Private = dto.Private
	user.Role = dto.Role
	user.Verified = dto.Verify
	user.ReceiveNotifications = dto.ReceiveNotifications
	user.Biography = dto.Biography
	user.Taggable = dto.Taggable
	user.Username = dto.Username
	user.Website = dto.Website
	user.Surname = dto.Surname

	return &user
}

//register user
func (handler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("creating")

	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	// //var user data.User2
	var user services.RegisterRequest
	//var user data.User2
	err := json.NewDecoder(r.Body).Decode(&user) //dekodiran je dto sa stringom..sad hocu da parsiram string

	if err != nil {
		println(err)
		w.WriteHeader(http.StatusNoContent)
		return
	}

	DateOfBirthTime := parseString(user.DateOfBirth)
	newUser := createUserFromDTO(user, DateOfBirthTime)
	fmt.Println(user)
	existsByUsername := handler.Service.Repo.UserExistsByUsername(newUser.Username)
	existsByEmail := handler.Service.Repo.UserExistsByEmail(newUser.Email)

	if existsByEmail || existsByUsername {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newUser.SetPassword(newUser.Password)
	err = handler.Service.CreateUser(newUser)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	fmt.Println("Created.")
	w.WriteHeader(http.StatusCreated)

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

	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	token := r.Header.Get("Authorization")
	splitToken := strings.Split(token, "Bearer ")
	token = splitToken[1]

	var updateUserRequest services.UpdateUserRequest
	err := json.NewDecoder(r.Body).Decode(&updateUserRequest) //ovde se nalazi token sa informacijama
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	tknStr := token
	tokenObj := &data.Token{}

	tkn, err := jwt.ParseWithClaims(tknStr, tokenObj, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusConflict)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	fmt.Printf("%x\n", tokenObj.UserID)
	err = handler.Service.UpdateEditUser(&updateUserRequest, tokenObj.UserID) //ovde saljem update User request
	if err != nil {
		fmt.Println(err)

		w.WriteHeader(http.StatusAlreadyReported)
	}
	fmt.Println("Updated.")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
