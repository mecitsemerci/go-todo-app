package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mecitsemerci/clean-go-todo-api/app/api"
	"github.com/mecitsemerci/clean-go-todo-api/app/infra/config"
	_ "github.com/mecitsemerci/clean-go-todo-api/docs"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func Init() *gin.Engine {
	router := gin.Default()
	apiRouteGroup := router.Group("/api")

	// Init config
	config.Init()

	api.AddControllers(apiRouteGroup)

	//Middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// TODO: Add Error Handler

	//Swagger Configuration
	router.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, ""))

	return router
}
