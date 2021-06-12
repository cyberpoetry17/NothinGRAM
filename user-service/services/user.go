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

//nema zabranu unosa istog imejla i korisnickog imena
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

func (service *UserService) UserExistsByEmail(email string) (bool, error) {
	exists := service.Repo.UserExistsByEmail(email)
	return exists, nil
}

func (service *UserService) UserExistsByUsername(username string) (bool, error) {
	exists := service.Repo.UserExistsByEmail(username)
	return exists, nil
}

func (service *UserService) FindOneByEmailAndPassword(email, password string) map[string]interface{} {
	user := &data.User2{}

	if err := service.Repo.Database.Where("Email = ?", email).First(user).Error; err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Email address not found"}
		return resp
	}
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials. Please try again"}
		return resp
	}
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

	var resp = map[string]interface{}{"status": false, "message": "logged in"}
	resp["token"] = tokenString //Store the token in the response
	resp["user"] = user
	return resp
}
