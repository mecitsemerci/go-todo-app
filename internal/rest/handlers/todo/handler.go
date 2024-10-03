package todo

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/mecitsemerci/go-todo-app/internal/rest/resultor"
)

type Handler struct {
	todoService TodoService
	logger      *zap.Logger
}

func NewHandler(todoService TodoService, logger *zap.Logger) *Handler {
	return &Handler{
		todoService: todoService,
		logger:      logger,
	}
}

func (h *Handler) GetAll(c *fiber.Ctx) error {
	tasks, err := h.todoService.GetAll(c.UserContext())

	if err != nil {
		h.logger.Error("Failed to get all tasks", zap.Error(err))
		return resultor.UnprocessableEntity(c, err)
	}

	return resultor.Ok(c, tasks)
}

func (h *Handler) Find(c *fiber.Ctx) error {
	input := GetTaskInput{TaskID: c.Query("task_id")}

	if err := input.Validate(); err != nil {
		h.logger.Error("Invalid taskID", zap.Error(err))
		return resultor.BadRequest(c, err)
	}

	task, err := h.todoService.Find(c.UserContext(), input)

	if err != nil {
		h.logger.Error("Failed to find task", zap.Error(err))
		return resultor.NotFound(c, err)
	}

	return resultor.Ok(c, task)
}

func (h *Handler) Create(c *fiber.Ctx) error {
	createInput := new(CreateTaskInput)

	if err := c.BodyParser(createInput); err != nil {
		h.logger.Error("Failed to parse request body", zap.Error(err))
		return resultor.BadRequest(c, err)
	}

	output, err := h.todoService.Create(c.UserContext(), *createInput)

	if err != nil {
		h.logger.Error("Failed to create task", zap.Error(err))
		return resultor.UnprocessableEntity(c, err)
	}

	return resultor.Ok(c, output)
}

func (h *Handler) Update(c *fiber.Ctx) error {
	input := GetTaskInput{TaskID: c.Params("task_id")}

	if err := input.Validate(); err != nil {
		h.logger.Error("Invalid task ID", zap.Error(err))
		return resultor.BadRequest(c, err)
	}

	updateInput := new(UpdateTaskInput)

	if err := c.BodyParser(updateInput); err != nil {
		h.logger.Error("Failed to parse request body", zap.Error(err))
		return resultor.BadRequest(c, err)
	}

	updateInput.TaskID = input.TaskID

	err := h.todoService.Update(c.UserContext(), *updateInput)

	if err != nil {
		h.logger.Error("Failed to update task", zap.Error(err))
		return resultor.UnprocessableEntity(c, err)
	}

	return resultor.Ok(c, "")
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	input := DeleteTaskInput{TaskID: c.Params("task_id")}

	if err := input.Validate(); err != nil {
		h.logger.Error("Invalid task ID", zap.Error(err))
		return resultor.BadRequest(c, err)
	}

	err := h.todoService.Delete(c.UserContext(), input)

	if err != nil {
		h.logger.Error("Failed to delete task", zap.Error(err))
		return resultor.UnprocessableEntity(c, err)
	}

	return resultor.Ok(c, "")
}
