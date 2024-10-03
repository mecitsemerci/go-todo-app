package auth

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/mecitsemerci/go-todo-app/internal/rest/resultor"
	"github.com/mecitsemerci/go-todo-app/pkg/identity"
)

const(
	otelName = "Handler"
	NoCache = "no-cache"
)

type Handler struct {
	authService AuthService
	logger      *zap.Logger
}

func NewHandler(authService AuthService, logger *zap.Logger) *Handler {
	return &Handler{
		authService: authService,
		logger:      logger,
	}
}

func (h *Handler) Login(c *fiber.Ctx) error {
	loginInput := new(LoginInput)

	if err := c.BodyParser(loginInput); err != nil {
		h.logger.Error("Failed to parse request body", zap.Error(err))
		return resultor.BadRequest(c, err)
	}

	output, err := h.authService.Login(c.Context(), *loginInput)

	if err != nil {
		h.logger.Error("Failed to login", zap.Error(err))
		return resultor.Unauthorized(c, err)
	}
	c.Response().Header.Set(fiber.HeaderCacheControl, NoCache)

	return resultor.Ok(c, output)
}

func (h *Handler) Register(c *fiber.Ctx) error {
	createAccountInput := new(CreateAccountInput)

	if err := c.BodyParser(createAccountInput); err != nil {
		h.logger.Error("Failed to parse request body", zap.Error(err))
		return resultor.BadRequest(c, err)
	}

	output, err := h.authService.CreateUser(c.Context(), *createAccountInput)

	if err != nil {
		h.logger.Error("Failed to create user", zap.Error(err))
		return resultor.UnprocessableEntity(c, err)
	}

	return resultor.Created(c, output)
}

func (h *Handler) GetProfile(c *fiber.Ctx) error {

	currentUser := identity.GetCurrentUserFromContext(c.UserContext())

	if currentUser == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid user context",
		})
	}

	if currentUser.Claims == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid user context",
		})
	}

	output, err := h.authService.GetProfile(c.Context(), GetAccountInput{
		ID: currentUser.Claims.UserID,
	})

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(output)

}
