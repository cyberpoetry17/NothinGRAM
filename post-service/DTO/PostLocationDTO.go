package DTO

import (
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/google/uuid"
)

type PostLocationDTO struct {
	Location data.Location
	PostId uuid.UUID
}