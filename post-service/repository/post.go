package repository

import (
	"fmt"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
)

type PostRepo struct {
	Database *gorm.DB
}

func (repo *PostRepo) CreatePost(post *data.Post) (error,uuid.UUID) {
	var location data.Location
	if post.UserID.String() =="00000000-0000-0000-0000-000000000000"{
		return errors.New("User id required"),post.ID
	}
	if post.LocationID == uuid.Nil {
		repo.Database.Find(&location).Where("country = dumb")
		post.LocationID = location.IDLoc
	}
	result := repo.Database.Create(post)
	if result.Error == nil {
		return result.Error,post.ID
	}
	fmt.Println(result.RowsAffected)
	return nil,post.ID //sta s ovim nilom
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

func (repo *PostRepo) GetTagsForPost(postid string) ([]string,error){
	var posts []data.Post
	var frontList []string
	result := repo.Database.Preload("Tags").Find(&posts)
	for _,element := range posts{
		if element.ID.String() == postid && element.Tags != nil{
			for _,el := range element.Tags{
				frontList = append(frontList,el.TagName)
			}

		}
	}
	return frontList,result.Error
}

func (repo *PostRepo) GetNonPrivatePosts() []data.Post{
	var posts []data.Post
	var frontList []data.Post
	posts = repo.GetAll()
	for _,element := range posts{
		if element.Private == false{
			frontList = append(frontList,element)
		}
	}
	return frontList
}

func (repo *PostRepo) GetNonPrivatePostsForUser(id string) []data.Post{
	var posts []data.Post
	var frontList []data.Post
	posts = repo.GetPostsByUserID(id)
	for _,element := range posts{
		if element.Private == false{
			frontList = append(frontList,element)
		}
	}
	return frontList
}

func (repo *PostRepo) GetPostsByUserID(id string) []data.Post{
	var posts []data.Post
	var frontList []data.Post
	repo.Database.
		Preload("Tags").
		Preload("Comments").
		Preload("Likes").
		Preload("Dislikes").
		Find(&posts)
	for _,element := range posts{
		if element.UserID.String() == id{
			frontList = append(frontList, element)
		}
	}
	return frontList
}

func (repo *PostRepo) GetPostByPostID(id string) data.Post{
	var posts []data.Post
	var frontList data.Post
	repo.Database.
		Preload("Tags").
		Preload("Comments").
		Preload("Likes").
		Preload("Dislikes").
		Find(&posts)
	for _,element := range posts{
		if element.ID.String() == id{
			frontList = element
		}
	}
	return frontList
}


func(repo *PostRepo) GetUsernameByPostUserID(userid string) string {
	var backString string
	var posts = repo.GetAll()
	for _,element := range posts{
		if element.UserID.String() == userid{
			response, _ := http.Get("http://localhost:8004/username/{"+userid+"}")
			backString,_:=ioutil.ReadAll(response.Body)
			return string(backString)
		}
	}
	return backString
}

func (repo *PostRepo) AddTagToPost(tag *data.Tag,postId uuid.UUID) error{
	for _, element := range repo.GetAll(){
		if(element.ID == postId){
			log.Println(tag.ID.String())
			//element.Tags = append(element.Tags, *tag)
			//repo.Database.Model(&data.Post{}).Association("Tags").Append(tag)
			////return nil
			tag.Posts = append(tag.Posts, element)
			err :=  repo.Database.Save(&tag).Error
			//err := repo.Database.Save(&element).Error	//ovo radi ali kreira novi tag
			//repo.Database.Model(&element).Association("Tags").Append(&tag)
			//err:=repo.Database.Session(&gorm.Session{FullSaveAssociations: true}).Save(element).Error
			//err := repo.Database.Raw("INSERT INTO posts_tags (tag_id,post_id) VALUES (?,?)",tag.ID,element.ID).Error

			return err
		}
	}
	return nil
}

func (repo *PostRepo) AddLocationToPost(location data.Location,postId uuid.UUID) error{
	for _, element := range repo.GetAll(){
		if element.ID == postId {
			//location.Posts = append(location.Posts,repo.GetPostById(postId))
			element.LocationID = location.IDLoc
			//repo.Database.Model(&data.Post{}).Association("Tags").Append(tag)
			////return nil
			err := repo.Database.Save(&element).Error

			//err:=repo.Database.Session(&gorm.Session{FullSaveAssociations: true}).Save(element).Error
			//err := repo.Database.Raw("INSERT INTO posts_locations (location_id_loc,post_id) VALUES (?,?)",location.IDLoc.String(),element.ID.String()).Error
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
