package identity

type contextKey string

const (
	CurrentUserKey = contextKey("currentUser")
)

type CurrentUser struct {
	Claims *Claims `json:"claims,omitempty"`
}

func (c *CurrentUser) Role() string {
	return c.Claims.Role
}

func (c *CurrentUser) Email() string {
	return c.Claims.Email
}

func (c *CurrentUser) UserID() string {
	return c.Claims.UserID
}
