package data

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

//Token struct declaration
type Token struct {
	UserID   uuid.UUID
	Username string
	Email    string
	*jwt.StandardClaims
}
