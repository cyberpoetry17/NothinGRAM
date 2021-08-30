package repository

import (
	"fmt"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
)

type LocationRepo struct {
	Database *gorm.DB
}

func (repo LocationRepo) CreateLocation(location *data.Location) (error,uuid.UUID) {
	for _,el := range repo.GetAll(){
		if(strings.ToLower(el.Country) == strings.ToLower(location.Country) && strings.ToLower(el.City)== strings.ToLower(location.City) && strings.ToLower(el.Address)== strings.ToLower(location.Address)){
			return  nil, el.IDLoc
		}
	}
	result := repo.Database.Create(location)
	if(result.Error != nil){
		return result.Error,location.IDLoc
	}
	fmt.Println(result.RowsAffected)
	return nil,location.IDLoc
}

func (repo LocationRepo) GetLocationForPost (postId string) *data.Location{
	var location data.Location
	repo.Database.Preload("Posts",&location)
	repo.Database.Where("Id = ?",postId).Find(&location)
	return &location
}

func (repo LocationRepo) RemoveLocation (location *data.Location) error{
	return repo.Database.Delete(location).Error
}

func (repo *LocationRepo) GetAll() []data.Location{
	var locations []data.Location
	repo.Database.Preload("Posts").Find(&locations)
	return locations
}

func (repo *LocationRepo) GetPostByLocation(locationId string) []data.Post{
	var posts []data.Post
	var locationFound = repo.GetAll()
	for _,element := range locationFound{
		if element.IDLoc.String() ==  locationId{
			for _,el := range element.Posts{
				posts = append(posts, el)
			}
		}
	}
	return posts
}

func (repo *LocationRepo) FilterPublicMaterialByLocation(locationId string) []data.Post{
	var media []data.Post
	var backList []data.Post
	//var frontList []data.Post
	media = repo.GetPostByLocation(locationId)
	for _,element := range media{
		if element.Private == false{
			backList = append(backList,element)
		}
	}
	//for _,element := range media{				prosirenje funkcije za kad se ubaci user
	//	if element.UserID.isPublic(){
	//		append(frontList, element)
	//	}
	//}
	for _,el := range backList{
		fmt.Println("Results: " + el.ID.String())
	}
	return backList//frontList
}