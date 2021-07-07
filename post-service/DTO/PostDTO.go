package DTO

import (
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/google/uuid"
	"time"
)

type PostDTO struct{
	ID                   uuid.UUID `json:"ID"`
	Description			 string `json:"description"`
	PicturePath			 string `json:"picpath"`
	UserID				 uuid.UUID `json:"userid"`
	Timestamp			 time.Time `json:"timestamp"`
	Comments			 []data.Comment `json:"Comments"`
	LocationID			 uuid.UUID `json:"LocationID"`
	Media				[]data.Media `json:"Media"`
	Private				 bool	`json:"private"`
	City				string `json:"city"`
	Country 				string `json:"country"`
	Address 			string `json:"address"`
	Tags			[]string `json:"tags"`
}
