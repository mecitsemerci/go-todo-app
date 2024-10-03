package identity

import "context"

func GetCurrentUserFromContext(ctx context.Context) *CurrentUser {
	currentUser, ok := ctx.Value(CurrentUserKey).(*CurrentUser)

	if !ok {
		return nil
	}
	return currentUser
}

func SetCurrentUserInContext(ctx context.Context, currentUser *CurrentUser) context.Context {
	return context.WithValue(ctx, CurrentUserKey, currentUser)
}
