package services

import (
	"fmt"
	"time"

	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *repository.UserRepo
}

type UpdateUserRequest struct {
	ID                   uuid.UUID   `json:"id"`
	Name                 string      `json:"name"`
	Surname              string      `  json:"surname"`
	Email                string      `  json:"email"`
	Username             string      ` json:"username"`
	Password             string      ` json:"password"`
	DateOfBirth          string      ` json:"date"`
	Gender               data.Gender ` json:"gender"`
	PhoneNumber          string      ` json:"phone"`
	Biography            string      ` json:"bio"`
	Website              string      ` json:"web"`
	Role                 data.Role   ` json:"role"`
	Private              bool        `json:"private"`
	Taggable             bool        `  json:"taggable"`
	ReceiveNotifications bool        `json:"notifications"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (service *UserService) CreateUser(user *data.User2) error {
	service.Repo.CreateUser(user)
	return nil
}

func (service *UserService) UserExists(userId string) (bool, error) {
	id, err := uuid.Parse(userId)
	if err != nil {
		print(err)
		return false, err
	}
	exists := service.Repo.UserExists(id)
	return exists, nil
}

func (service *UserService) GetUserById(id uuid.UUID) (*data.User2, error) {
	user, err := service.Repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *UserService) LoginUser(r *LoginRequest) map[string]interface{} {
	user := &data.User2{}
	//nadje i kastuje
	err := service.Repo.Database.Where("Email = ?", r.Email).First(user).Error
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Email address not found"}
		return resp
	}
	//setuje vreme
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()
	//poredi hesirane passworde da vidi da li su jednaki
	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.Password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials. Please try again"}
		return resp
	}
	//pravi novi token sa informacijama
	tk := &data.Token{
		UserID:   user.ID,
		Username: user.Username,
		Email:    user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}

	var resp = map[string]interface{}{"status": true, "message": "logged in"}
	resp["token"] = tokenString //Store the token in the response
	resp["user"] = user
	return resp
}

func checkIfStringIsValid(toCheck string) bool {
	return toCheck != ""
}

func (service *UserService) UpdateEditUser(r *UpdateUserRequest) error {
	user, error := service.Repo.GetById(r.ID)
	if error != nil {
		fmt.Println("ovo ovde je greska")
		return error
	}
	if checkIfStringIsValid(r.Name) {
		user.Name = r.Name
	}
	if checkIfStringIsValid(r.Surname) {
		user.Surname = r.Surname
	}
	if checkIfStringIsValid(r.Biography) {
		user.Biography = r.Biography
	}
	if checkIfStringIsValid(r.Website) {
		user.Website = r.Website
	}
	if checkIfStringIsValid(r.PhoneNumber) {
		user.PhoneNumber = r.PhoneNumber
	}
	if checkIfStringIsValid(r.Email) && r.Email != user.Email {
		fmt.Println("uslo u imejl")
		emailTaken := service.Repo.UserExistsByEmail(r.Email)
		error2 := fmt.Errorf("email already taken")
		if emailTaken {
			return error2
		}
		user.Email = r.Email
	}
	if checkIfStringIsValid(r.Username) && (r.Username != user.Username) {
		fmt.Println("uslo u username")
		//
		service.Repo.UserExistsByUsername(r.Username)
		// error2 := fmt.Errorf("username already taken")
		// if usernameTaken {
		// 	return error2
		// }
		user.Username = r.Username
	}
	if checkIfStringIsValid(r.Password) && r.Password != user.Password {

		user.SetPassword(r.Password)
	}

	user.Gender = r.Gender
	user.Private = r.Private
	user.ReceiveNotifications = r.ReceiveNotifications
	user.Taggable = r.Taggable
	user.Role = r.Role
	user.DateOfBirth = r.DateOfBirth
	errorUpdatingUser := service.Repo.Database.Save(&user).Error
	if errorUpdatingUser != nil {
		return errorUpdatingUser
	}
	return nil
}
