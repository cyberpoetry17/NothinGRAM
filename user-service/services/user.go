package services

import (
	"fmt"
	"time"

	"github.com/cyberpoetry17/NothinGRAM/UserAPI/DTO"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo              *repository.UserRepo
	RepoFollower      *repository.FollowerRepo
	RepoCloseFollower *repository.CloseFollowerRepository
	MutedRepo         *repository.MutedRepo
	BlockedRepo       *repository.BlockedRepo
}

type UpdateUserRequest struct {
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
	//Verify               bool        `json:"verify"`
}

type RegisterRequest struct {
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
	Verify               bool        `json:"verify"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (service *UserService) CreateUser(user *data.User2) error {
	//println(user.DateOfBirth)
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

func (service *UserService) GetUserById(ID uuid.UUID) (*data.User2, error) {
	user, err := service.Repo.GetById(ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *UserService) GetUserByUsernameForProfile(id string) *data.User2 {
	return service.Repo.GetUserByUsernameForProfile(id)
}

func (service *UserService) GetPublicUserIds() []string {
	return service.Repo.GetPublicUserIds()
}

func (service *UserService) DeleteProfile(id string) bool{
	service.BlockedRepo.DeleteBlocksForUser(id)
	service.MutedRepo.DeleteMutesForUser(id)
	service.RepoFollower.DeleteFollowersForUser(id)
	service.RepoCloseFollower.DeleteCloseFollowersForUser(id)
	service.Repo.DeleteUserById(id)
	return true
}

func (service *UserService) GetUserIdByUsernameForProfile(id string) DTO.UserUsernameAndPrivateDTO {
	return service.Repo.GetUserIdByUsernameForProfile(id)
}

func (service *UserService) GetUsernameById(id uuid.UUID) string {
	return service.Repo.GetUsernameById(id)
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

	expiresAt := time.Now().Local().Add(time.Minute * 10000)
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
		Role:     int64(user.Role),
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
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
	resp["expirationDate"] = expiresAt
	return resp
}

func checkIfStringIsValid(toCheck string) bool {
	return toCheck != ""
}

func (service *UserService) ChangePassword(r *LoginRequest) error {
	user, error := service.Repo.GetByEmail(r.Email)
	if error != nil {
		fmt.Println("ovo ovde je greska")
		return error
	}
	if checkIfStringIsValid(r.Password) && r.Password != user.Password {

		user.SetPassword(r.Password)
	}
	errorUpdatingUser := service.Repo.Database.Save(&user).Error
	if errorUpdatingUser != nil {
		return errorUpdatingUser
	}
	return nil

}

func (service *UserService) ParseString(info string) time.Time {
	t, err := time.Parse("2006-01-02T15:04:05.000Z", info)
	if err != nil {
		println("Time parsing not supported!")
	}
	return t
}

func (service *UserService) UpdateEditUser(r *UpdateUserRequest, ID uuid.UUID) error {
	user, error := service.Repo.GetById(ID)
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
		service.Repo.UserExistsByUsername(r.Username)
		user.Username = r.Username
	}
	if checkIfStringIsValid(r.Password) && r.Password != user.Password {
		user.SetPassword(r.Password)
	}
	user.Gender = r.Gender
	user.Private = r.Private
	user.ReceiveNotifications = r.ReceiveNotifications
	println(r.ReceiveNotifications)
	user.Taggable = r.Taggable
	user.Role = r.Role
	user.DateOfBirth = service.ParseString(r.DateOfBirth)
	errorUpdatingUser := service.Repo.Database.Save(&user).Error
	if errorUpdatingUser != nil {
		return errorUpdatingUser
	}
	return nil
}

func (service *UserService) GetUserProfilePrivacy(ID uuid.UUID) (bool, error) {
	user, err := service.Repo.GetById(ID)
	return user.Private, err
}

func (service *UserService) GetAllById(id uuid.UUID) []string {
	user, err := service.Repo.GetById(id)
	followers := user.Followers
	allUsers := service.Repo.GetAll()

	if err != nil {
		return nil
	}

	var userFollowers []string
	for _, elementFollower := range followers {
		for _, elementUser := range allUsers {
			if elementUser.ID == elementFollower.IDFollower {
				userFollowers = append(userFollowers, elementUser.Username)
			}
		}
	}
	return userFollowers
}

func (service *UserService) SetCloseFollowersToUser(list []string, id uuid.UUID) []DTO.UserDTO {
	var followers []DTO.UserDTO
	if list == nil {
		return nil

	}
	for _, username := range list { //GetUserIdByUsernameForProfile
		var user DTO.UserDTO
		usernameDto, err := service.Repo.GetUserIdByUsername(username)
		fmt.Println(usernameDto.Private)
		user.UserId = usernameDto.ID
		if err != nil {
			return nil
		}
		followers = append(followers, user)
	}
	return followers
}

func (service *UserService) GetAllCloseFollowersById(id uuid.UUID) []string {
	user, err := service.Repo.GetById(id)
	followers := user.CloseFollowers
	allUsers := service.Repo.GetAll()

	if err != nil {
		return nil
	}

	var userFollowers []string
	for _, elementFollower := range followers {
		for _, elementUser := range allUsers {
			if elementUser.ID == elementFollower.IDCloseFollower {
				userFollowers = append(userFollowers, elementUser.Username)
			}
		}
	}
	return userFollowers
}
