//+build wireinject

package wired

import (
	"github.com/google/wire"
	"github.com/mecitsemerci/go-todo-app/internal/api/handler"
	"github.com/mecitsemerci/go-todo-app/internal/core/services"
	"github.com/mecitsemerci/go-todo-app/internal/pkg/mongodb"
)

// MongoDB Service Dependencies
var TodoRepositorySetByMongo = wire.NewSet(mongodb.ProvideTodoRepository, mongodb.ProvideMongoClient)
var TodoServiceSetByMongo = wire.NewSet(services.ProvideTodoService, TodoRepositorySetByMongo, mongodb.ProvideIDGenerator)

func InitializeTodoHandler() (handler.TodoHandler, error) {

	wire.Build(handler.ProvideTodoHandler, TodoServiceSetByMongo)

	return handler.TodoHandler{}, nil
}

func InitializeHealthHandler() handler.HealthHandler {
	return handler.HealthHandler{}
}
