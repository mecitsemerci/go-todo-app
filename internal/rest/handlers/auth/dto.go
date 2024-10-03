package auth

import "time"

type CreateAccountInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateAccountOutput struct {
	ID string `json:"id"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginOutput struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type LogoutInput struct {
	Token string `json:"token"`
}

type LogoutOutput struct {
	Success bool `json:"success"`
}

type GetAccountInput struct {
	ID string `json:"id"`
}

type GetAccountOutput struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type UpdateAccountInput struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UpdateAccountOutput struct {
	Success bool `json:"success"`
}

type UserProfileOutput struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	LastLogin time.Time `json:"last_login"`
}
