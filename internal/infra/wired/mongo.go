//go:build wireinject
// +build wireinject

package wired

import (
	"github.com/google/wire"

	handler2 "github.com/mecitsemerci/go-todo-app/internal/app/api/handler"
	"github.com/mecitsemerci/go-todo-app/internal/core/service"
	"github.com/mecitsemerci/go-todo-app/internal/infra/mongodb"
)

// MongoDB Service Dependencies
var TodoRepositorySetByMongo = wire.NewSet(mongodb.ProvideTodoRepository, mongodb.ProvideMongoClient)
var TodoServiceSetByMongo = wire.NewSet(service.ProvideTodoService, TodoRepositorySetByMongo, mongodb.ProvideIDGenerator)

func InitializeTodoHandler() (handler2.TodoHandler, error) {

	wire.Build(handler2.ProvideTodoHandler, TodoServiceSetByMongo)

	return handler2.TodoHandler{}, nil
}

func InitializeHealthHandler() handler2.HealthHandler {
	return handler2.HealthHandler{}
}
