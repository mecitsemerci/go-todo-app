//+build wireinject

package wired

import (
	"github.com/google/wire"
	"github.com/mecitsemerci/go-todo-app/internal/api/handler"
	"github.com/mecitsemerci/go-todo-app/internal/core/services"
	"github.com/mecitsemerci/go-todo-app/internal/pkg/redisdb"
)

// Redis Service Dependencies
var TodoRepositorySetByRedis = wire.NewSet(redisdb.ProvideTodoRepository, redisdb.ProvideRedisClient)
var TodoServiceSetByRedis = wire.NewSet(services.ProvideTodoService, TodoRepositorySetByRedis, redisdb.ProvideIDGenerator)

func InitializeTodoController() (handler.TodoHandler, error) {

	wire.Build(handler.ProvideTodoHandler, TodoServiceSetByRedis)

	return handler.TodoHandler{}, nil
}

func InitializeHealthController() handler.HealthController {
	return handler.HealthController{}
}
