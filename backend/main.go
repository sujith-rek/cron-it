package main

import (
	"cronbackend/controllers"
	"cronbackend/db"
	"cronbackend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	userController controllers.UserController
	userRouter     routes.UserRouter

	ScheduleController controllers.ScheduleController
	ScheduleRouter     routes.ScheduleRouter

	CheckController controllers.CheckingController
	CheckRouter     routes.CheckRouter

	server *gin.Engine
)

func init() {
	config, err := db.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	db.ConnectDB(&config)

	userController = *controllers.NewUserController(db.DB)
	userRouter = *routes.NewUserRouter(&userController)

	ScheduleController = *controllers.NewScheduleController(db.DB)
	ScheduleRouter = *routes.NewScheduleRouter(&ScheduleController)

	CheckController = *controllers.NewCheckingController(db.DB)
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
