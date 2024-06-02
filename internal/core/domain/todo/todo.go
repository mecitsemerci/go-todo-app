package todo

import (
	"time"

	"github.com/mecitsemerci/go-todo-app/internal/core/enum"
)

// Todo represents information
type Todo struct {
	ID          ID
	ProjectID   ID
	Title       string
	Description string
	Completed   bool
	Priority    enum.PriorityLevel
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// DefaultTodo is Default Todo
var DefaultTodo Todo
