package service

import (
	"context"

	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"

	"github.com/mecitsemerci/go-todo-app/config"
	"github.com/mecitsemerci/go-todo-app/internal/domain/auth"
	dto "github.com/mecitsemerci/go-todo-app/internal/rest/handlers/auth"
	"github.com/mecitsemerci/go-todo-app/pkg/identity"
)

type Auth struct {
	userRepo auth.UserRepository
	logger   *zap.Logger
}

func NewAuth(userRepo auth.UserRepository, logger *zap.Logger) *Auth {
	return &Auth{
		userRepo: userRepo,
		logger:   logger,
	}
}

func (a *Auth) CreateUser(ctx context.Context, input dto.CreateAccountInput) (*dto.CreateAccountOutput, error) {
	spanCtx, span := otel.Tracer(otelName).Start(ctx, "AuthService.CreateAccount")
	defer span.End()

	isExist, err := a.userRepo.IsEmailTaken(spanCtx, auth.Email(input.Email))

	if err != nil {
		return nil, errors.Wrap(err, "is email taken failed")
	}

	if isExist {
		return nil, errors.New("email already taken")
	}

	user, err := auth.NewUser(auth.Email(input.Email), input.Password)

	if err != nil {
		return nil, errors.Wrap(err, "new user failed")
	}

	err = a.userRepo.Insert(spanCtx, user)

	if err != nil {
		return nil, errors.Wrap(err, "insert failed")
	}

	return &dto.CreateAccountOutput{
		ID: user.ID.String(),
	}, nil
}

func (a *Auth) GetProfile(ctx context.Context, input dto.GetAccountInput) (*dto.UserProfileOutput, error) {
	spanCtx, span := otel.Tracer(otelName).Start(ctx, "AuthService.GetProfile")
	defer span.End()

	user, err := a.userRepo.GetByID(spanCtx, auth.UserID(input.ID))

	if err != nil {
		return nil, errors.Wrap(err, "find by id failed")
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	return &dto.UserProfileOutput{
		ID:        user.ID.String(),
		Email:     user.Email.String(),
		Role:      user.Role.String(),
		CreatedAt: user.CreatedAt,
		LastLogin: user.LastLogin,
	}, nil
}

func (a *Auth) UpdateProfile(ctx context.Context, input dto.UpdateAccountInput) (*dto.UpdateAccountOutput, error) {
	spanCtx, span := otel.Tracer(otelName).Start(ctx, "AuthService.UpdateProfile")
	defer span.End()

	user, err := a.userRepo.GetByID(spanCtx, auth.UserID(input.ID))

	if err != nil {
		return nil, errors.Wrap(err, "find by id failed")
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	err = user.Update(auth.Email(input.Email), input.Password)

	if err != nil {
		return nil, errors.Wrap(err, "update email or password failed")
	}

	if input.Role != "" && user.IsAdmin() {
		user.UpdateRole(auth.Role(input.Role))
	}

	err = a.userRepo.Update(spanCtx, user)

	if err != nil {
		return nil, errors.Wrap(err, "update failed")
	}

	return &dto.UpdateAccountOutput{
		Success: true,
	}, nil
}

func (a *Auth) Login(ctx context.Context, input dto.LoginInput) (*dto.LoginOutput, error) {
	spanCtx, span := otel.Tracer(otelName).Start(ctx, "AuthService.Login")
	defer span.End()

	user, err := a.userRepo.FindByEmail(spanCtx, auth.Email(input.Email))

	if err != nil {
		return nil, errors.Wrap(err, "find by email failed")
	}

	if user == nil {
		return nil, errors.New("invalid email or password")
	}

	if err := user.Password.Verify(input.Password); err != nil {
		return nil, errors.Wrap(err, "invalid email or password")
	}

	claims := identity.Claims{
		UserID: user.ID.String(),
		Role:   user.Role.String(),
		Email:  user.Email.String(),
	}
	a.logger.Debug("claims", zap.Any("claims", claims))
	// Set expire
	claims.SetExpire(config.AppConfig.JWTExpiresSec)
	a.logger.Debug("exp", zap.Any("exp", claims.ExpiresAt))

	// Create token
	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims.MapClaims())

	// Generate encoded token and send it as response.
	accessToken, err := token.SignedString([]byte(config.AppConfig.JWTSecret))

	if err != nil {
		return nil, errors.Wrap(err, "token signing failed")
	}

	return &dto.LoginOutput{
		AccessToken: accessToken,
		TokenType:   identity.BearerSchema,
		ExpiresIn:   config.AppConfig.JWTExpiresSec,
	}, nil
}

func (a *Auth) Logout(ctx context.Context, input dto.LogoutInput) (*dto.LogoutOutput, error) {
	return nil, nil
}

func (a *Auth) Verify(ctx context.Context, email auth.Email, password string) error {
	user, err := a.userRepo.FindByEmail(ctx, email)

	if err != nil {
		return err
	}

	if user == nil {
		return auth.ErrUserNotFound
	}

	return user.Password.Verify(password)
}
