package services

import (
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/DTO"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/repository"
	"github.com/google/uuid"
	"time"
)

type PostService struct {
	PostRepo *repository.PostRepo
	TagRepo  *repository.TagRepo
	LikeRepo *repository.LikeRepo
	DislikeRepo *repository.DislikeRepo
	MediaRepo	*repository.MediaRepo
}

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
		media.PostId = idPost
		media.Type = data.Picture
		media.Link = el
		service.MediaRepo.CreateMedia(&media)
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

func (service *PostService) GetNonPrivatePostsForUser(id string) []data.Post{
	return service.PostRepo.GetNonPrivatePostsForUser(id)
}

func (service *PostService) GetPostsByUserID(id string) []data.Post{
	return service.PostRepo.GetPostsByUserID(id)
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