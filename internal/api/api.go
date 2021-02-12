package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mecitsemerci/go-todo-app/internal/wired"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func AddControllers(apiRouteGroup *gin.RouterGroup) {

	groupV1 := apiRouteGroup.Group("/v1")
	{
		todoController := wired.InitializeTodoController()
		todoController.Register(groupV1)
	}

	healthController := wired.InitializeHealthController()
	healthController.Register(apiRouteGroup)

}

func Setup() *gin.Engine {
	router := gin.Default()
	apiRouteGroup := router.Group("/api")

	AddControllers(apiRouteGroup)

	//Swagger Configuration
	router.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, ""))

	return router
}
