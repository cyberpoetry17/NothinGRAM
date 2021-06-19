package repository

import (
	"fmt"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostRepo struct {
	Database *gorm.DB
}

func (repo *PostRepo) CreatePost(post *data.Post) error {
	var location data.Location
	if post.LocationID == uuid.Nil {
		repo.Database.Find(&location).Where("country = dumb")
		post.LocationID = location.IDLoc
	}
		result := repo.Database.Create(post)
		if result.Error == nil {
			return result.Error
		}
		fmt.Println(result.RowsAffected)
		return nil //sta s ovim nilom
}

func (repo *PostRepo) PostExists(desc string) bool {
	var count int64
	repo.Database.Where("picpath", desc).Find(&data.Post{}).Count(&count)
	return count != 0
}

func (repo *PostRepo) EditPost(post *data.Post) error  {
	return repo.Database.Save(post).Error
}

func (repo *PostRepo) GetAll() []data.Post{
	var posts []data.Post
	repo.Database.
		Preload("Tags").
		Preload("Comments").
		Preload("Likes").
		Preload("Dislikes").
		Find(&posts)
	return posts
}

func (repo *PostRepo) AddTagToPost(tag data.Tag,postId uuid.UUID) error{
	for _, element := range repo.GetAll(){
		if(element.ID == postId){
			element.Tags = append(element.Tags, tag)
			//repo.Database.Model(&data.Post{}).Association("Tags").Append(tag)
			////return nil
			//err := repo.Database.Save(&element).Error	//ovo radi ali kreira novi tag

			//err:=repo.Database.Session(&gorm.Session{FullSaveAssociations: true}).Save(element).Error
			err := repo.Database.Raw("INSERT INTO posts_tags (tag_id,post_id) VALUES (?,?)",tag.ID.String(),element.ID.String()).Error

			return err
		}
	}
	return nil
}

// func (r *UserRepo) SaveUser(user *data.User2) *data.User2 {
// 	//databaseError := map[string]string{}
// 	err := r.database.Debug().Create(&user).Error //CREATE prosledjenog usera,vratim mapu gresaka i usera
// 	if err != nil {
// 		//If the email is already taken
// 		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
// 			//databaseError["email_taken"] = "email already taken"
// 			return nil
// 		}
// 		//any other db error
// 		//databaseError["db_error"] = "database error"
// 		return nil
// 	}
// 	return user
// }

// //GET USER
// func (r *UserRepo) GetUser(id uint64) *data.User2 {
// 	var user data.User2
// 	err := r.database.Debug().Where("id = ?", id).Take(&user).Error //GET USER BY ID vraca usera i nista za gresku
// 	if err != nil {
// 		return nil
// 	}
// 	// if gorm.IsRecordNotFoundError(err) {
// 	// 	return nil, errors.New("user not found")
// 	// }
// 	return &user
// }

// //GET LIST OF USERS
// // func (r *UserRepo) GetAllUsers() data.Users {
// // 	var users data.Users
// // 	err := r.database.Debug().Find(&users).Error
// // 	fmt.Println("repozitorijum nije greska")
// // 	if err != nil {
// // 		fmt.Println("repozitorijum greska")
// // 		return nil
// // 	}
// // 	return users
// // }

// func (r *UserRepo) GetAllUsers() (data.Users, error) {
// 	var users data.Users
// 	err := r.database.Debug().Find(&users).Error
// 	if err != nil {

// 		return nil, err
// 	}
// 	sampleErr := errors.New("good")
// 	return users, sampleErr
// }

// func (r *UserRepo) GetUserByEmailAndPassword(u *data.User2) *data.User2 {
// 	var user data.User2 //user kojeg trazim
// 	//databaseError := map[string]string{}
// 	errr := r.database.Debug().Where("email = ?", u.Email).Take(&user).Error

// 	// if gorm.ErrRecordNotFound(errr) {
// 	// 	databaseError["no_user_found"] = "Error : user not found!"
// 	// 	return nil, databaseError
// 	// }
// 	if errr != nil {
// 		//	databaseError["database_error"] = "database error"
// 		return nil
// 	}

// 	//password checking needed
// 	return &user

// }

// //GET USER BY EMAIL
// func (ur *UserRepo) OneByEmail(email string) *data.User2 {
// 	var u data.User2
// 	errr := ur.database.Debug().Where("email = ?", u.Email).Take(&u).Error
// 	if errr != nil {

// 		return nil
// 	}
// 	return &u
// }
