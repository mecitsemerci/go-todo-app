//+build wireinject

package wired

import (
	"github.com/google/wire"
	"github.com/mecitsemerci/clean-go-todo-api/app/api/controller"
	v1 "github.com/mecitsemerci/clean-go-todo-api/app/api/controller/v1"
	"github.com/mecitsemerci/clean-go-todo-api/app/core"
	"github.com/mecitsemerci/clean-go-todo-api/app/infra/repository"
)

var todoServiceSet = wire.NewSet(core.ProvideTodoService, repository.ProvideTodoRepository, repository.ProvideDbContext)

func InitializeTodoControllerV1() v1.TodoController {

	wire.Build(controller.ProvideTodoController, todoServiceSet)

	return v1.TodoController{}
}

func InitializeHealthController() controller.HealthController {
	wire.Bind(controller.ProvideHealthController, repository.ProvideDbContext)
	return controller.HealthController{}
}
