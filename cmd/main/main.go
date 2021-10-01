package main

import (
	"github.com/gin-gonic/gin"
	"homecredithack-task/pkg/database"
	"homecredithack-task/pkg/handler"
	"homecredithack-task/pkg/repository"
	"homecredithack-task/pkg/service"
)

func main() {
	app := gin.Default()
	defer app.Run(":8081") //8080

	repos := repository.NewRepository(database.Connect())
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	handlers.InitRoutes(app)

	database.Connect()

}
