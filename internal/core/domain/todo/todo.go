package todo

import (
	"github.com/mecitsemerci/go-todo-app/internal/core/domain"
	"github.com/mecitsemerci/go-todo-app/internal/core/enum"
	"time"
)

type Todo struct {
	ID          domain.ID
	Title       string
	Description string
	Completed   bool
	Priority    enum.PriorityLevel
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

