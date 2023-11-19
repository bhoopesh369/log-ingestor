package router

import (
	"github.com/labstack/echo/v4"
)

func InitRoutes(server *echo.Echo) {
	server.GET("/ping", ping)
	logRouter := server.Group("/logs")
	injestLogs(logRouter)
}
