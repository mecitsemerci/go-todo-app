package core

import "github.com/mecitsemerci/clean-go-todo-api/app/core/services"

func ProvideTodoService() *services.TodoService {
	service := new(services.TodoService)
	return service
}
