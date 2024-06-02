package redisdb

import (
	"context"
	"fmt"
	"github.com/mecitsemerci/go-todo-app/config"
	"github.com/mecitsemerci/go-todo-app/internal/core/domain/todo"
	"time"

	"github.com/go-redis/redis/v8"
)

// ProvideIDGenerator provides IDGenerator
func ProvideIDGenerator() todo.IDGenerator {
	return NewIDGenerator()
}

// ProvideRedisClient provides mongo client
func ProvideRedisClient() (*redis.Client, error) {
	//Set Options
	opts := redis.Options{
		Addr:        config.RedisConfig.RedisURL,
		DB:          config.RedisConfig.RedisDb,
		PoolSize:    config.RedisConfig.RedisMaxPoolSize,
		DialTimeout: time.Second * time.Duration(config.RedisConfig.RedisConnectionTimeout),
	}

	rdb := redis.NewClient(&opts)

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("redis : there is no connection. %v", err)
	}

	return rdb, nil
}

// ProvideTodoRepository provides redis adapter
func ProvideTodoRepository(client *redis.Client) todo.TodoRepository {
	return NewTodoAdapter(client)
}
