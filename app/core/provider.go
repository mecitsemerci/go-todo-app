package core

import (
	"github.com/mecitsemerci/clean-go-todo-api/app/core/domain/todo"
	"github.com/mecitsemerci/clean-go-todo-api/app/core/services"
)

func ProvideTodoService(todoRepository todo.ITodoRepository) services.ITodoService {
	return services.NewTodoService(todoRepository)
}