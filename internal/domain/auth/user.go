package auth

import (
	"time"
)

type User struct {
	ID        UserID
	Email     Email
	Password  PasswordHash
	Role      Role
	CreatedAt time.Time
	UpdatedAt time.Time
	LastLogin time.Time
}

func NewUser(email Email, password string) (*User, error) {
	if !email.IsValid() {
		return nil, ErrInvalidEmail
	}
	passwordHash, err := NewPasswordHash(password)
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return &User{
		ID:        NewUserID(),
		Email:     email,
		Password:  passwordHash,
		Role:      Member,
		CreatedAt: now,
		UpdatedAt: now,
		LastLogin: now,
	}, nil
}

func (a *User) Update(email Email, password string) error {
	if !email.IsValid() {
		return ErrInvalidEmail
	}

	passwordHash, err := NewPasswordHash(password)
	if err != nil {
		return err
	}
	a.Email = email
	a.Password = passwordHash
	a.UpdatedAt = time.Now().UTC()
	return nil
}

func (a *User) Equals(other *User) bool {
	return a.ID.Equals(other.ID)
}

func (a *User) IsAdmin() bool {
	return a.Role == Admin
}

func (a *User) UpdateRole(role Role) {
	a.Role = role
	a.UpdatedAt = time.Now().UTC()
}

func (a *User) VerifyPassword(password string) error {
	return a.Password.Verify(password)
}

func (a *User) Authenticate(password string) bool {
	return a.Password.Verify(password) == nil
}

func (a *User) UpdateLastLogin() {
	a.LastLogin = time.Now().UTC()
}
