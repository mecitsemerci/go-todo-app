package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mecitsemerci/clean-go-todo-api/app/infra/wired"
)

func AddControllers(apiRouteGroup *gin.RouterGroup) {

	groupV1 := apiRouteGroup.Group("/v1")
	{
		todoControllerV1 := wired.InitializeTodoControllerV1()
		todoControllerV1.Register(groupV1)
	}

	healthController := wired.InitializeHealthController()
	healthController.Register(apiRouteGroup)
}
