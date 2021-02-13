package mongodb

import (
	"context"
	"github.com/mecitsemerci/go-todo-app/internal/config"
	"github.com/mecitsemerci/go-todo-app/internal/core/interfaces"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func ProvideIdGenerator() interfaces.IDGenerator {
	return NewIdGenerator()
}

func ProvideMongoClient() *mongo.Client {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(config.MongoDbUrl))
	if err != nil {
		log.Printf("provide mongodb: %s", err.Error())
	}
	return client
}

func ProvideTodoRepository(client *mongo.Client) interfaces.TodoRepository {
	return NewTodoAdapter(client, config.MongoDbTodoDbName)
}
