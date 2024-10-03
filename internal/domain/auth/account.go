package auth

type Account struct {
	Email    Email
	Session  string
	Sessions []UserSession
}
