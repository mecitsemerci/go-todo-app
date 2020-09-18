package controller

import (
	v1 "github.com/mecitsemerci/clean-go-todo-api/app/api/controller/v1"
	"github.com/mecitsemerci/clean-go-todo-api/app/core/services"
)

func ProvideController(todoService services.ITodoService) v1.TodoController {
	return v1.NewTodoController(todoService)
}