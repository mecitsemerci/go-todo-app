package interfaces

import "github.com/mecitsemerci/go-todo-app/internal/core/domain"

//IDGenerator represents generic domain ID generator
type IDGenerator interface {
	NewID() domain.ID
	IDFromString(str string) (domain.ID, error)
}
