package main

import (
	"github.com/bhoopesh369/log-injestor/config"
	"github.com/bhoopesh369/log-injestor/router"
	"github.com/bhoopesh369/log-injestor/services"
	"github.com/bhoopesh369/log-injestor/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	server := echo.New()

	utils.InitLogger(server)
	server.Use(middleware.CORS())
	server.Use(middleware.Recover())

	config.InitConfig()
	config.ConnectDB()
	config.MigrateDB()

	router.InitRoutes(server)

	services.SummaService()

	services.ProducerService()
	services.ConsumerService()

	server.Logger.Fatal(server.Start(":" + config.ServerPort))
}
