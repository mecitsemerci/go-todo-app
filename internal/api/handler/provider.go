package handler

import (
	"github.com/mecitsemerci/go-todo-app/internal/core/interfaces"
)

//ProvideTodoHandler returns new TodoHandler instance for wire
func ProvideTodoHandler(todoService interfaces.TodoService) TodoHandler {
	return NewTodoHandler(todoService)
}
