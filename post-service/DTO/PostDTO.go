package DTO

import (
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/google/uuid"
	"time"
)

type PostDTO struct{
	ID                   uuid.UUID ` json:"ID"` //mozda ovo ne treba?
	Description			 string `json:"description"`
	Likes				 []data.Like	`json:"likes"`
	Dislikes			 []data.Dislike	`json:"dislikes"`
	UserID				 uuid.UUID `json:"userid"`
	Timestamp			 time.Time `json:"timestamp"`
	Tags 				 []data.Tag `json:"Tags"`
	Comments			 []data.Comment `json:"Comments"`
	LocationID			 uuid.UUID `json:"LocationID"`
	Media				[]data.Media `json:"Media"`
	Private				 bool	`json:"private"`
	ImgPaths			[]string `json:"ImgPaths"`
	VideoPath			string	`json:"VideoPath"`
}
