package service

import (
	"github.com/mecitsemerci/go-todo-app/internal/core/domain/todo"
)

// ProvideTodoService provides TodoService according to interface
func ProvideTodoService(todoRepository TodoRepository, idGenerator IDGenerator[todo.ID]) *TodoService {
	return NewTodoService(todoRepository, idGenerator)
}
