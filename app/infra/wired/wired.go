//+build wireinject

package wired

import (
	"github.com/google/wire"
	"github.com/mecitsemerci/clean-go-todo-api/app/api/controller"
	"github.com/mecitsemerci/clean-go-todo-api/app/core"
	"github.com/mecitsemerci/clean-go-todo-api/app/infra/repository"
)

var todoServiceSet = wire.NewSet(core.ProvideTodoService, repository.ProvideTodoRepository, repository.ProvideDbContext)

func InitializeTodoControllerV1() controller.TodoControllerV1 {

	wire.Build(controller.ProvideTodoController, todoServiceSet)

	return controller.TodoControllerV1{}
}

func InitializeHealthController() controller.HealthController {
	wire.Bind(controller.ProvideHealthController, repository.ProvideDbContext)
	return controller.HealthController{}
}
