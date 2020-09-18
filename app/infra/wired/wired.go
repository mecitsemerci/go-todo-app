//+build wireinject

package wired

import (
	"github.com/google/wire"
	"github.com/mecitsemerci/clean-go-todo-api/app/api/controller"
	v1 "github.com/mecitsemerci/clean-go-todo-api/app/api/controller/v1"
	"github.com/mecitsemerci/clean-go-todo-api/app/core"
	"github.com/mecitsemerci/clean-go-todo-api/app/infra/adapter"
)

var todoServiceSet = wire.NewSet(core.ProvideTodoService, adapter.ProvideTodoRepository, adapter.ProvideDbContext)

func InitializeTodoControllerV1() v1.TodoController {

	wire.Build(controller.ProvideController, todoServiceSet)

	return v1.TodoController{}
}
