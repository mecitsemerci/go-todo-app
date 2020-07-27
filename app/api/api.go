package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mecitsemerci/clean-go-todo-api/app/api/controller"
	v1 "github.com/mecitsemerci/clean-go-todo-api/app/api/controller/v1"
)



func AddControllers(apiRouteGroup *gin.RouterGroup) {

	groupV1 := apiRouteGroup.Group("/v1")
	{
		new(v1.TodoController).RegisterRoutes(groupV1)
	}

	new(controller.HealthController).RegisterRoutes(apiRouteGroup)
}
