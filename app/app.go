package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mecitsemerci/clean-go-todo-api/app/api"
	"github.com/mecitsemerci/clean-go-todo-api/app/infra/adapter/mementodb"
	_ "github.com/mecitsemerci/clean-go-todo-api/docs"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
)

func Init() {
	router := gin.Default()
	apiRouteGroup := router.Group("/api")

	api.AddControllers(apiRouteGroup)

	//Init Db
	db := mementodb.DbContext{}
	db.Seed()

	//Swagger Configuration
	router.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, ""))

	// Run app
	if err := router.Run(":8080"); err != nil {
		log.Println(err)
		return
	}
}
