package controller

import (
	v1 "github.com/mecitsemerci/clean-go-todo-api/app/api/controller/v1"
	"github.com/mecitsemerci/clean-go-todo-api/app/core/services"
	"github.com/mecitsemerci/clean-go-todo-api/app/infra/repository/mongodb"
)

func ProvideTodoController(todoService services.ITodoService) v1.TodoController {
	return v1.NewTodoController(todoService)
}
func ProvideHealthController(dbContext mongodb.DbContext) HealthController  {
	return NewHealthController(dbContext)
}