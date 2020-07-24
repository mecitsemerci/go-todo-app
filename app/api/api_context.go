package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mecitsemerci/clean-go-todo-api/app/api/controllers"
)

func AddControllers(apiRouteGroup *gin.RouterGroup)  {
	controllers.RegisterControllers(apiRouteGroup)
}
