package routes

import (
	"cronbackend/controllers"
	"cronbackend/middleware"
	"github.com/gin-gonic/gin"
)

type ScheduleRouter struct {
	ScheduleController *controllers.ScheduleController
}

func NewScheduleRouter(sc *controllers.ScheduleController) *ScheduleRouter {
	return &ScheduleRouter{ScheduleController: sc}
}

func (sr *ScheduleRouter) RegisterRoutes(router *gin.RouterGroup) {
	rg := router.Group("/schedules")
	rg.Use(middleware.UnwrapUserToken())

	// return meow as string response
	rg.GET("/", func (c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "meow",
		})
	})

	rg.POST("/create", sr.ScheduleController.CreateJobSchedule)
	rg.GET("/print", sr.ScheduleController.PrintCluster)
	
}