package main

import (
	"cronbackend/controller"
	"cronbackend/db"
	"cronbackend/routes"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"net/http"

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

func recoverOnRestart(){
	// call localhost:8080/recover to recover all jobs
	response, err := http.Get("http://localhost:8080/recover")
	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode != 200 {
		log.Fatal("Failed to recover jobs")
	}

	log.Println("Recovered all jobs", response.Status)
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

	recoverOnRestart()

	log.Fatal(server.Run(":" + config.ServerPort))


}
