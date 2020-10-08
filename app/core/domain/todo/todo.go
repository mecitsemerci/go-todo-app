package todo

import (
	"github.com/mecitsemerci/clean-go-todo-api/app/core/domain"
	"time"
)

type Todo struct {
	Id          domain.ID
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
