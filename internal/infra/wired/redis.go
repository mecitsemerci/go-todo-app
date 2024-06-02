//go:build wireinject
// +build wireinject

package wired

import (
	"github.com/google/wire"

	handler2 "github.com/mecitsemerci/go-todo-app/internal/app/api/handler"
	"github.com/mecitsemerci/go-todo-app/internal/core/service"
	"github.com/mecitsemerci/go-todo-app/internal/infra/redisdb"
)

// Redis Service Dependencies
var TodoRepositorySetByRedis = wire.NewSet(redisdb.ProvideTodoRepository, redisdb.ProvideRedisClient)
var TodoServiceSetByRedis = wire.NewSet(service.ProvideTodoService, TodoRepositorySetByRedis, redisdb.ProvideIDGenerator)

func InitializeTodoHandler() (handler2.TodoHandler, error) {

	wire.Build(handler2.ProvideTodoHandler, TodoServiceSetByRedis)

	return handler2.TodoHandler{}, nil
}

func InitializeHealthHandler() handler2.HealthHandler {
	return handler2.HealthHandler{}
}
