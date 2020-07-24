package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mecitsemerci/clean-go-todo-api/app/api"
	"github.com/mecitsemerci/clean-go-todo-api/app/infra/adapters/mementodb"
	"log"
)

func Init() {
	router := gin.Default()
	apiRouteGroup := router.Group("/api")

	api.AddControllers(apiRouteGroup)

	//Init Db
	db := mementodb.DbContext{}
	db.Seed()

	// Run app
	if err := router.Run(":8080"); err != nil {
		log.Println(err)
		return
	}
}
