package mongodb

import (
	"context"
	"time"

	"github.com/mecitsemerci/go-todo-app/internal/config"
	"github.com/mecitsemerci/go-todo-app/internal/core/interfaces"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//ProvideIDGenerator provides IDGenerator
func ProvideIDGenerator() interfaces.IDGenerator {
	return NewIDGenerator()
}

//ProvideMongoClient provides mongo client
func ProvideMongoClient() (*mongo.Client, error) {
	//Set Options
	opts := options.Client().ApplyURI(config.MongoConfig.MongoURL)
	maxPoolSize := config.MongoConfig.MongoMaxPoolSize
	connTimeout := time.Duration(config.MongoConfig.MongoConnectionTimeout) * time.Second
	opts.MaxPoolSize = &maxPoolSize
	opts.ConnectTimeout = &connTimeout
	return mongo.Connect(context.Background(), opts)
}

//ProvideTodoRepository provides mongodb adapter
func ProvideTodoRepository(client *mongo.Client) interfaces.TodoRepository {
	return NewTodoAdapter(client, config.MongoConfig.MongoTodoDbName)
}
