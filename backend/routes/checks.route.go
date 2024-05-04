package routes

import (
	"cronbackend/controller"
	"cronbackend/middleware"
	"github.com/gin-gonic/gin"
)

type CheckRouter struct {
	CheckController *controller.CheckingController
}

func NewCheckRouter(cc *controller.CheckingController) *CheckRouter {
	return &CheckRouter{CheckController: cc}
}

func (cr *CheckRouter) RegisterRoutes(router *gin.RouterGroup) {
	rg := router.Group("/checks")
	rg.Use(middleware.UnwrapUserToken())
}
