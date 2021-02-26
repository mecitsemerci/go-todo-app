package services

import (
	"github.com/mecitsemerci/go-todo-app/internal/core/interfaces"
)

//ProvideTodoService provides TodoService according to interface
func ProvideTodoService(todoRepository interfaces.TodoRepository,
	idGenerator interfaces.IDGenerator) interfaces.TodoService {
	return NewTodoService(todoRepository, idGenerator)
}
