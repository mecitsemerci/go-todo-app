package auth

import "context"

type AuthService interface {
	CreateUser(context.Context, CreateAccountInput) (*CreateAccountOutput, error)
	GetProfile(context.Context, GetAccountInput) (*UserProfileOutput, error)
	UpdateProfile(context.Context, UpdateAccountInput) (*UpdateAccountOutput, error)
	Login(context.Context, LoginInput) (*LoginOutput, error)
	Logout(context.Context, LogoutInput) (*LogoutOutput, error)
}
