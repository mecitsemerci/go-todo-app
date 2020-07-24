package controllers

import (
	"github.com/gin-gonic/gin"
)

func RegisterControllers(apiRouteGroup *gin.RouterGroup) {
	new(TodoController).RegisterRoutes(apiRouteGroup)
	new(HealthController).RegisterRoutes(apiRouteGroup)
}
