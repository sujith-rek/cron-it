package routes

import (
	"cronbackend/controllers"
	"cronbackend/middleware"
	"github.com/gin-gonic/gin"
)

type CheckRouter struct {
	CheckController *controllers.CheckingController
}

func NewCheckRouter(cc *controllers.CheckingController) *CheckRouter {
	return &CheckRouter{CheckController: cc}
}

func (cr *CheckRouter) RegisterRoutes(router *gin.RouterGroup) {
	rg := router.Group("/checks")
	rg.Use(middleware.UnwrapUserToken())

	rg.POST("/ping/:userID/:pingID", cr.CheckController.Ping)
}
