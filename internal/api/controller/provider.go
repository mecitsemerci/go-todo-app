package controller

import (
	"github.com/mecitsemerci/go-todo-app/internal/core/interfaces"
)

func ProvideTodoController(todoService interfaces.TodoService) TodoController {
	return NewTodoController(todoService)
}
func ProvideHealthController() HealthController {
	return NewHealthController()
}
