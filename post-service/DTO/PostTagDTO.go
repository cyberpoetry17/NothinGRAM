package DTO

import (
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/google/uuid"
)

type PostTagDTO struct {
	Tag data.Tag
	PostId uuid.UUID
}