package routes

import (
	"cronbackend/controllers"
	"cronbackend/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	UserController *controllers.UserController
}

func NewUserRouter(uc *controllers.UserController) *UserRouter {
	return &UserRouter{UserController: uc}
}

func (ur *UserRouter) RegisterRoutes(router *gin.RouterGroup) {

	rg := router.Group("/users")

	rg.POST("/register", ur.UserController.Register)
	rg.POST("/login", ur.UserController.Login)
	rg.POST("/refresh", ur.UserController.RefreshToken)
	rg.POST("/logout", middleware.UnwrapUserToken(), ur.UserController.Logout)
}
