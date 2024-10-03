package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"

	"github.com/mecitsemerci/go-todo-app/config"
	"github.com/mecitsemerci/go-todo-app/internal/rest/resultor"
	"github.com/mecitsemerci/go-todo-app/pkg/identity"
)

func NewAuthorize(logger *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {

		authHeaderVal := c.Get(fiber.HeaderAuthorization)

		if authHeaderVal == "" {
			logger.Debug("Missing authorization header")
			return resultor.Unauthorized(c, ErrAuthorizeHeaderMissing)
		}

		tokenStr := extractBearerToken(authHeaderVal)
		if tokenStr == "" {
			logger.Debug("Invalid authorization token", zap.String("token", tokenStr))
			return resultor.Unauthorized(c, ErrTokenInvalid)
		}

		claims, err := parseToken(tokenStr)

		if err != nil {
			logger.Error("Failed to parse token", zap.Error(err))
			return resultor.Unauthorized(c, err)
		}

		currentUser := identity.CurrentUser{
			Claims: claims,
		}

		ctx := identity.SetCurrentUserInContext(c.Context(), &currentUser)

		c.SetUserContext(ctx)

		return c.Next()
	}
}

func extractBearerToken(header string) string {
	parts := strings.Split(header, " ")
	if len(parts) != 2 || parts[0] != identity.BearerSchema {
		return ""
	}
	return parts[1]
}

func parseToken(tokenString string) (*identity.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &identity.Claims{}, func(token *jwt.Token) (any, error) {
		return []byte(config.AppConfig.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*identity.Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrTokenNotValidYet
}
