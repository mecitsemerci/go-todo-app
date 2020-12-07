package core

import (
	"github.com/mecitsemerci/clean-go-todo-api/app/core/domain/todo"
	"github.com/mecitsemerci/clean-go-todo-api/app/core/services"
)

func ProvideTodoService(todoRepository todo.Repository) services.ITodoService {
	return services.NewTodoService(todoRepository)
}