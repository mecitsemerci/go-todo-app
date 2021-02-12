package services

import (
	"github.com/mecitsemerci/go-todo-app/internal/core/interfaces"
)

func ProvideTodoService(todoRepository interfaces.TodoRepository,
	idGenerator interfaces.IDGenerator) interfaces.TodoService {
	return NewTodoService(todoRepository, idGenerator)
}
