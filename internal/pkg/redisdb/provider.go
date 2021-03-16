package redisdb

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/mecitsemerci/go-todo-app/internal/config"
	"github.com/mecitsemerci/go-todo-app/internal/core/interfaces"
)

//ProvideIDGenerator provides IDGenerator
func ProvideIDGenerator() interfaces.IDGenerator {
	return NewIDGenerator()
}

//ProvideRedisClient provides mongo client
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

//ProvideTodoRepository provides redis adapter
func ProvideTodoRepository(client *redis.Client) interfaces.TodoRepository {
	return NewTodoAdapter(client)
}
