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
	router.POST(("/refresh"), ur.UserController.RefreshToken)
	router.POST(("/logout"), middleware.UnwrapUserToken(), ur.UserController.Logout)
}