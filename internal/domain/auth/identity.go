package auth

import (
	"github.com/google/uuid"
)

type UserID string

func (u UserID) String() string {
	return string(u)
}

func (u UserID) Equals(other UserID) bool {
	return u.String() == other.String()
}

func NewUserID() UserID {
	return UserID(uuid.New().String())
}
