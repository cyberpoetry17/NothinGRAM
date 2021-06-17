package data

import "github.com/google/uuid"

type Follower struct {
	UserID     uuid.UUID
	FollowerID uuid.UUID
	Accepted   bool
}
