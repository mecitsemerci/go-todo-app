package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/mecitsemerci/go-todo-app/config"
	"github.com/mecitsemerci/go-todo-app/internal/core/domain/todo"
	"github.com/mecitsemerci/go-todo-app/internal/core/service"
)

// ProvideIDGenerator provides IDGenerator
func ProvideIDGenerator() service.IDGenerator[todo.ID] {
	return NewIDGenerator()
}

// ProvideMongoClient provides mongo client
func ProvideMongoClient() (*mongo.Client, error) {
	//Set Options
	opts := options.Client().ApplyURI(config.MongoConfig.MongoURL)
	maxPoolSize := config.MongoConfig.MongoMaxPoolSize
	connTimeout := time.Duration(config.MongoConfig.MongoConnectionTimeout) * time.Second
	opts.MaxPoolSize = &maxPoolSize
	opts.ConnectTimeout = &connTimeout
	return mongo.Connect(context.Background(), opts)
}

// ProvideTodoRepository provides mongodb adapter
func ProvideTodoRepository(client *mongo.Client) service.TodoRepository {
	return NewTodoAdapter(client, config.MongoConfig.MongoTodoDbName)
}
