package identity

import (
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/mecitsemerci/go-todo-app/internal/domain/auth"
)
const (
	BearerSchema = "Bearer"
)

type Claims struct {
	UserID string `json:"user_id,omitempty"`
	Role   string `json:"role,omitempty"`
	Email  string `json:"email,omitempty"`
	jwt.StandardClaims
}

func (c *Claims) SetRole(role auth.Role) *Claims {
	c.Role = role.String()
	return c
}

func (c *Claims) SetUserID(userID auth.UserID) *Claims {
	c.UserID = userID.String()
	return c
}

func (c *Claims) SetEmail(email auth.Email) *Claims {
	c.Email = email.String()
	return c
}

func (c *Claims) SetExpire(seconds int) *Claims {
	c.ExpiresAt = time.Now().Add(time.Duration(seconds) * time.Second).Unix()
	return c
}

func (c *Claims) MapClaims() jwt.MapClaims {
	return jwt.MapClaims{
		"user_id": c.UserID,
		"role":    c.Role,
		"email":   c.Email,
		"exp":     c.ExpiresAt,
	}
}
