package middleware

import "errors"

var (
	ErrAuthorizeHeaderMissing = errors.New("authorization header missing")
	ErrTokenExpired           = errors.New("token expired")
	ErrTokenNotValidYet       = errors.New("token not valid yet")
	ErrTokenMalformed         = errors.New("malformed token")
	ErrTokenInvalid           = errors.New("invalid token")
)
