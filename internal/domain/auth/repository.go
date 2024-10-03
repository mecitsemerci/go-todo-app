package auth

import (
	"context"
)

type UserRepository interface {
	Insert(ctx context.Context, user *User) error
	FindByEmail(ctx context.Context, email Email) (*User, error)
	GetByID(ctx context.Context, userID UserID) (*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, userID UserID) error
	UpdateLastLogin(ctx context.Context, userID UserID) error
	IsEmailTaken(ctx context.Context, email Email) (bool, error)
	UpdatePassword(ctx context.Context, userID UserID, password string) error
}
