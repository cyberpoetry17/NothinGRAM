package DTO

import "github.com/google/uuid"

type UserUsernameAndPrivateDTO struct {
	UserId uuid.UUID
	Private bool
}