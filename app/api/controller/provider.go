package controller

import (
	"github.com/mecitsemerci/clean-go-todo-api/app/core/services"
	"github.com/mecitsemerci/clean-go-todo-api/app/infra/repository/mongodb"
)

func ProvideTodoController(todoService services.ITodoService) TodoControllerV1 {
	return NewTodoController(todoService)
}
func ProvideHealthController(dbContext mongodb.DbContext) HealthController  {
	return NewHealthController(dbContext)
}