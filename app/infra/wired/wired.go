//+build wireinject

package wired

import (
	v1 "github.com/mecitsemerci/clean-go-todo-api/app/api/controller/v1"
)

func InitializeTodo() (v1.TodoController, error)  {
	// TODO:
	return v1.TodoController{}
}

