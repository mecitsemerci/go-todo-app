package resultor

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Response[T any] struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
	Data    T      `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func Success[T any](c *fiber.Ctx, status int, message string, data T) error {
	response := Response[T]{
		Status:  status,
		Message: message,
		Data:    data,
	}

	return c.Status(status).JSON(response)
}

func Error[T error](c *fiber.Ctx, status int, message string, err T) error {
	response := Response[T]{
		Status:  status,
		Message: message,
		Error:   err.Error(),
	}

	return c.Status(status).JSON(response)
}

func Ok[T any](c *fiber.Ctx, data T) error {
	return Success(c, http.StatusOK, "", data)
}

func Created[T any](c *fiber.Ctx, data T) error {
	return Success(c, http.StatusCreated, "", data)
}

func BadRequest(c *fiber.Ctx, err error) error {
	return Error(c, http.StatusBadRequest, "", err)
}

func InternalServerError(c *fiber.Ctx, err error) error {
	return Error(c, http.StatusInternalServerError, "", err)
}

func Unauthorized(c *fiber.Ctx, err error) error {
	return Error(c, http.StatusUnauthorized, "", err)
}

func Forbidden(c *fiber.Ctx, err error) error {
	return Error(c, http.StatusForbidden, "", err)
}

func NotFound(c *fiber.Ctx, err error) error {
	return Error(c, http.StatusNotFound, "", err)
}

func Conflict(c *fiber.Ctx, err error) error {
	return Error(c, http.StatusConflict, "", err)
}

func UnprocessableEntity(c *fiber.Ctx, err error) error {
	return Error(c, http.StatusUnprocessableEntity, "", err)
}
