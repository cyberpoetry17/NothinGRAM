package DTO

import "github.com/google/uuid"

type StoryMediaDTO struct{
	MediaPath string	`json:"MediaPath"`
	UserId 	uuid.UUID 	`json:"UserId"`
}