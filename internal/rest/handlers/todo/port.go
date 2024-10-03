package todo

import (
	"context"
)

type TodoService interface {
	GetAll(ctx context.Context) (*GetTasksOutput, error)
	Find(ctx context.Context, input GetTaskInput) (*TaskOutput, error)
	Create(ctx context.Context, input CreateTaskInput) (*TaskOutput, error)
	Update(ctx context.Context, input UpdateTaskInput) error
	Delete(ctx context.Context, input DeleteTaskInput) error
}
