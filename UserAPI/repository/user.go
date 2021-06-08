package repository

import (
	"strings"

	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"gorm.io/gorm"
)

type UserRepo struct {
	database *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{db} //pravi novi repo
}

var _ IUserRepository = &UserRepo{}

func (r *UserRepo) SaveUser(user *data.User2) (*data.User2, map[string]string) {
	databaseError := map[string]string{}
	err := r.database.Debug().Create(&user).Error //CREATE prosledjenog usera,vratim mapu gresaka i usera
	if err != nil {
		//If the email is already taken
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			databaseError["email_taken"] = "email already taken"
			return nil, databaseError
		}
		//any other db error
		databaseError["db_error"] = "database error"
		return nil, databaseError
	}
	return user, nil
}

func (r *UserRepo) GetUser(id uint64) (*data.User2, error) {
	var user data.User2
	err := r.database.Debug().Where("id = ?", id).Take(&user).Error //GET USER BY ID vraca usera i nista za gresku
	if err != nil {
		return nil, err //ne vraca usera ali vraca gresku
	}
	// if gorm.IsRecordNotFoundError(err) {
	// 	return nil, errors.New("user not found")
	// }
	return &user, nil
}

func (r *UserRepo) GetAllUsers() (data.Users, error) {
	var users data.Users
	err := r.database.Debug().Find(&users).Error
	if err != nil {
		return nil, err
	}
	// if gorm.IsRecordNotFoundError(err) {
	// 	return nil, errors.New("user not found")
	// }
	return users, nil
}

func (r *UserRepo) GetUserByEmailAndPassword(u *data.User2) (*data.User2, map[string]string) {
	var user data.User2 //user kojeg trazim
	databaseError := map[string]string{}
	errr := r.database.Debug().Where("email = ?", u.Email).Take(&user).Error

	// if gorm.ErrRecordNotFound(errr) {
	// 	databaseError["no_user_found"] = "Error : user not found!"
	// 	return nil, databaseError
	// }
	if errr != nil {
		databaseError["database_error"] = "database error"
		return nil, databaseError
	}

	//password checking needed
	return &user, nil

}
