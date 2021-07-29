package services

import (
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/DTO"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/repository"
	"github.com/google/uuid"
	"strings"
	"time"
)

type PostService struct {
	PostRepo *repository.PostRepo
	TagRepo  *repository.TagRepo
	LikeRepo *repository.LikeRepo
	DislikeRepo *repository.DislikeRepo
	MediaRepo	*repository.MediaRepo
	LocationRepo *repository.LocationRepo
}
var extensions =[]string {"mp4","mov","avi","wmv","m4a"}
const Pi = 3

//verovatno treba da vrati neku vrednost
func (service *PostService) CreatePost(postDto *DTO.PostDTO) error {
	var postToAdd data.Post
	postToAdd.Tags = postDto.Tags
	postToAdd.LocationID = postDto.LocationID
	postToAdd.UserID = postDto.UserID
	postToAdd.Private = postDto.Private
	postToAdd.Description = postDto.Description
	postToAdd.Timestamp = time.Now()

	err,idPost := service.PostRepo.CreatePost(&postToAdd)
	if err != nil {
		return err
	}
	for _,el := range postDto.ImgPaths{
		var media data.Media
		media.PostId = &idPost
		media.Type = data.Picture
		media.Link = el
		for _, e := range extensions {
			split := strings.Split(media.Link, "?")
			if(len(split)==0){
				continue
			}
			if(strings.HasSuffix(split[0], e)){
				media.Type= data.Video
				break
			}
		}

		err=service.MediaRepo.CreateMedia(&media)
		if err !=nil{
			return err
		}
	}

	return err
}

func (service *PostService) PostExists(desc string) (bool, error) {
	id := desc
	exists := service.PostRepo.PostExists(id)
	return exists, nil
}

func (service *PostService) AddTagToPost(tag data.Tag,postId uuid.UUID) error{
	var tagToAdd data.Tag
	service.TagRepo.Database.Where("id = ?",tag.ID).Find(&tagToAdd)
	return service.PostRepo.AddTagToPost(&tagToAdd,postId)
}

func (service *PostService) AddLocationToPost(location data.Location,postId uuid.UUID) error{
	return service.PostRepo.AddLocationToPost(location,postId)
}
func (service *PostService) GetAllPosts() []data.Post{
	return service.PostRepo.GetAll()
}

func (service *PostService) GetNonPrivatePosts() []data.Post{
	return service.PostRepo.GetNonPrivatePosts()
}

func (service *PostService) GetNonPrivatePostsForUser(id string) ([]data.Post,error){
	posts,err := service.PostRepo.GetNonPrivatePostsForUser(id)
	if err != nil{
		return nil,err
	}

	return posts,err
}

func (service *PostService) GetPostsByUserID(id string) []data.Post{
	return service.PostRepo.GetPostsByUserID(id)
}

func (service *PostService) GetPostsByLocation(id string) []data.Post{
	var frontList []data.Post
	var locations []data.Location
	service.LocationRepo.Database.Find(&locations)
	for _,element := range locations{
		if strings.ToLower(element.Country) == strings.ToLower(id) || strings.ToLower(element.City) == strings.ToLower(id) || strings.ToLower(element.Address) == strings.ToLower(id) {
			frontList = append(frontList,service.PostRepo.GetPostsByLocationId(element.IDLoc.String())...)
		}
	}
	return frontList
}

func (service *PostService) GetPostsByTags(id string) []data.Post{
	var frontList []data.Post
	var tag []data.Tag
	service.TagRepo.Database.Find(&tag)
	for _,element := range tag{
		if strings.ToLower(element.TagName) == strings.ToLower(id) {
			frontList = append(frontList,service.PostRepo.GetPostsByTagId(element.ID.String())...)
		}
	}
	return frontList
}

func (service *PostService) GetUsernameByPostUserID(userid string) string{
	return service.PostRepo.GetUsernameByPostUserID(userid)
}

func (service *PostService) GetTagsForPost(postid string) ([]string,error){
	tags,err := service.PostRepo.GetTagsForPost(postid)
	if err != nil{
		return nil,err
	}
	return tags,err
}

func (service *PostService) GetLikedByUser(userid string) []data.Post{
	var likedPosts []data.Like
	var frontList []data.Post
	service.LikeRepo.Database.Find(&likedPosts)
	for _,element := range likedPosts{
		if element.UserId == userid {
			frontList = append(frontList, service.PostRepo.GetPostByPostID(element.PostId.String()))
		}
	}
	return frontList
}

func (service *PostService) GetDislikedByUser(userid string) []data.Post{
	var dislikedPosts []data.Dislike
	var frontList []data.Post
	service.DislikeRepo.Database.Find(&dislikedPosts)
	for _,element := range dislikedPosts{
		if element.UserId == userid {
			frontList = append(frontList, service.PostRepo.GetPostByPostID(element.PostId.String()))
		}
	}
	return frontList
}