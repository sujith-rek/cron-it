package main

import (
	"cronbackend/controller"
	"cronbackend/db"
	"cronbackend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	userController controller.UserController
	userRouter     routes.UserRouter

	ScheduleController controller.ScheduleController
	ScheduleRouter     routes.ScheduleRouter

	CheckController controller.CheckingController
	CheckRouter     routes.CheckRouter

	server *gin.Engine
)

func init() {
	config, err := db.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	db.ConnectDB(&config)

	userController = *controller.NewUserController(db.DB)
	userRouter = *routes.NewUserRouter(&userController)

	ScheduleController = *controller.NewScheduleController(db.DB)
	ScheduleRouter = *routes.NewScheduleRouter(&ScheduleController)

	CheckController = *controller.NewCheckingController(db.DB)
	CheckRouter = *routes.NewCheckRouter(&CheckController)

	server = gin.Default()

}

func main() {
	config, err := db.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")

	userRouter.RegisterRoutes(router)
	ScheduleRouter.RegisterRoutes(router)
	CheckRouter.RegisterRoutes(router)

	log.Fatal(server.Run(":" + config.ServerPort))

}
