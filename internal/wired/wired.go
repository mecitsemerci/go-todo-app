//+build wireinject

package wired

import (
	"github.com/google/wire"
	"github.com/mecitsemerci/go-todo-app/internal/api/handler"
	"github.com/mecitsemerci/go-todo-app/internal/core/services"
	"github.com/mecitsemerci/go-todo-app/internal/pkg/mongodb"
)

var TodoRepositorySet = wire.NewSet(mongodb.ProvideTodoRepository, mongodb.ProvideMongoClient)
var TodoServiceSet = wire.NewSet(services.ProvideTodoService, TodoRepositorySet, mongodb.ProvideIDGenerator)

func InitializeTodoController() (handler.TodoHandler, error) {
	wire.Build(handler.ProvideTodoHandler, TodoServiceSet)
	return handler.TodoHandler{}, nil
}

func InitializeHealthController() handler.HealthController {
	return handler.HealthController{}
}
