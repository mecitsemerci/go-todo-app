package auth

import "errors"

var (
	ErrInvalidEmail      = errors.New("invalid email address")
	ErrUserNotFound      = errors.New("user not found")
	ErrInvalidPassword   = errors.New("invalid password")
	ErrUserAlreadyExists = errors.New("user already exists")
)
