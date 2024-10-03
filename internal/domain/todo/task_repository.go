package todo

import (
	"context"

	"github.com/mecitsemerci/go-todo-app/internal/domain/auth"
)

type TaskRepository interface {
	GetAll(ctx context.Context) ([]*Task, error)
	GetAllByUser(ctx context.Context, userID auth.UserID) ([]*Task, error)
	GetByID(ctx context.Context, taskID TaskID) (*Task, error)
	Insert(ctx context.Context, task *Task) error
	Update(ctx context.Context, task *Task) error
	Delete(ctx context.Context, taskID TaskID) error
}
