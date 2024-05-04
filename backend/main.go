package main

import (
	"cronbackend/controller"
	"cronbackend/db"
	"cronbackend/routes"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

)

var (
	userController controller.UserController
	userRouter     routes.UserRouter
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

	log.Fatal(server.Run(":" + config.ServerPort))


}
