package todo

import (
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type CreateTaskInput struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Priority    int    `json:"priority" validate:"required"`
}

type CreateTaskOutput struct {
	TaskID string `json:"task_id"`
}

type UpdateTaskInput struct {
	TaskID      string `json:"task_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
	Completed   bool   `json:"completed"`
}

type DeleteTaskInput struct {
	TaskID string `json:"task_id" validate:"required"`
}

func (i *DeleteTaskInput) Validate() error {
	if i.TaskID == "" {
		return ErrTaskIDRequired
	}

	if _, err := uuid.Parse(i.TaskID); err != nil {
		return errors.Wrap(err, "invalid task ID")
	}

	return nil
}

type GetTaskInput struct {
	TaskID string `json:"task_id" validate:"required"`
}

func (i *GetTaskInput) Validate() error {
	if i.TaskID == "" {
		return ErrTaskIDRequired
	}

	if _, err := uuid.Parse(i.TaskID); err != nil {
		return errors.Wrap(err, "invalid task ID")
	}

	return nil
}

type TaskOutput struct {
	TaskID      string    `json:"task_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	Priority    int       `json:"priority"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetTasksOutput struct {
	Tasks      []*TaskOutput `json:"tasks,omitempty"`
	TotalCount int           `json:"total_count"`
	Page       int           `json:"page"`
}
