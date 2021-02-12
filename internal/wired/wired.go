//+build wireinject

package wired

import (
	"github.com/google/wire"
	"github.com/mecitsemerci/go-todo-app/internal/api/controller"
	"github.com/mecitsemerci/go-todo-app/internal/core/services"
	"github.com/mecitsemerci/go-todo-app/internal/pkg/mongodb"
)

var TodoRepositorySet = wire.NewSet(mongodb.ProvideTodoRepository, mongodb.ProvideMongoClient)
var TodoServiceSet = wire.NewSet(services.ProvideTodoService, TodoRepositorySet, mongodb.ProvideIdGenerator)

func InitializeTodoController() controller.TodoController {

	wire.Build(controller.ProvideTodoController, TodoServiceSet)

	return controller.TodoController{}
}

func InitializeHealthController() controller.HealthController {
	return controller.HealthController{}
}
