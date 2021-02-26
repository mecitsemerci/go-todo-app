package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mecitsemerci/go-todo-app/internal/pkg/tracer"
	"github.com/mecitsemerci/go-todo-app/internal/wired"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

//AppEngine application runtime
type AppEngine struct {
	Close func() error
	Run   func() error
}

func registerHandlers(apiRouteGroup *gin.RouterGroup) {

	groupV1 := apiRouteGroup.Group("/v1")
	{
		todoController := wired.InitializeTodoController()
		todoController.Register(groupV1)
	}

	healthController := wired.InitializeHealthController()
	healthController.Register(apiRouteGroup)

}

//Setup returns AppEngine
func Setup() *AppEngine {
	router := gin.Default()
	apiRouteGroup := router.Group("/api")

	registerHandlers(apiRouteGroup)

	//Opentracing configuration
	traceCloser := tracer.Init()

	//Swagger Configuration
	router.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, ""))

	return &AppEngine{
		Close: func() error {
			return traceCloser.Close()
		},
		Run: func() error {
			return router.Run()
		},
	}
}
