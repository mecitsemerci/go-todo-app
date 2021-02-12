package interfaces

import "github.com/mecitsemerci/go-todo-app/internal/core/domain"

type IDGenerator interface {
	NewID() domain.ID
	IDFromString(str string) (domain.ID, error)
}
