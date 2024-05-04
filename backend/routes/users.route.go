package routes

import (
	"cronbackend/controller"
	"cronbackend/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	UserController *controller.UserController
}

func NewUserRouter(uc *controller.UserController) *UserRouter {
	return &UserRouter{UserController: uc}
}

func (ur *UserRouter) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/register", ur.UserController.Register)
	router.POST("/login", ur.UserController.Login)
	router.GET(("/refresh"), ur.UserController.RefreshToken)
	router.GET(("/logout"), middleware.UnwrapUserToken(), ur.UserController.Logout)
}