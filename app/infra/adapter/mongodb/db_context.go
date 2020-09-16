package mongodb

import (
	"context"
	"github.com/mecitsemerci/clean-go-todo-api/app/infra/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type DbContext struct {
	Context        context.Context
	Client         *mongo.Client
	Database       *mongo.Database
	TodoCollection *mongo.Collection
}

func (dbContext *DbContext) init(timeout time.Duration) *DbContext {
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoDbUrl))
	if err != nil {
		log.Fatal("MongoDb connection is refused. Error: ", err)
	}
	database := client.Database(config.MongoDbName)
	dbContext.Context = ctx
	dbContext.Client = client
	dbContext.Database = database
	dbContext.TodoCollection = database.Collection("Todo")
	return dbContext
}
func (dbContext *DbContext) Connect() {
	dbContext.init(15*time.Second)
}

func (dbContext *DbContext) ConnectWithTimeout(timeout time.Duration) {
	dbContext.init(timeout)
}

func (dbContext *DbContext) Disconnect()  {
	dbContext.Client.Disconnect(dbContext.Context)
}

