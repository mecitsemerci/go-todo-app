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
func ProvideMongoClient() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.DBTimeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoURL))
	if err != nil {
		panic(err)
		//log.Printf("provide mongodb: %s", err.Error())
	}
	return client
}

//ProvideTodoRepository provides mongodb adapter
func ProvideTodoRepository(client *mongo.Client) interfaces.TodoRepository {
	return NewTodoAdapter(client, config.MongoTodoDbName)
}
