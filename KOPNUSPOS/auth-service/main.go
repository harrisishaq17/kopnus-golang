package main

import (
	"auth-service/config"
	"auth-service/controller"
	"auth-service/helpers"
	"auth-service/repository"
	"auth-service/service"
	"log"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	var e = echo.New()
	e.Use(middleware.CORS())
	e.Validator = &helpers.EchoValidator{Validator: validator.New()}

	config.LoadConfig()
	dbInit := config.InitDBConnection(config.GormConfig.DBHost, config.GormConfig.DBUser, config.GormConfig.DBPassword, config.GormConfig.DBName, config.GormConfig.DBPort)
	logrusInit := config.InitLogrus(config.LogrusConfig.LogLevel, config.LogrusConfig.GraylogHost, config.LogrusConfig.GraylogPort)

	// setup repositories
	usersRepo := repository.NewUserRepository(logrusInit)

	// setup services
	usersService := service.NewUserService(dbInit, logrusInit, *usersRepo)

	// setup controller
	usersController := controller.NewUserController(*usersService)

	usersController.UserRoutes(e)

	e.Logger.Fatal(e.Start(":3003"))
}
