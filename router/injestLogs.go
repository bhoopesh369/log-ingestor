package router

import (
	"github.com/bhoopesh369/log-injestor/controllers"
	"github.com/labstack/echo/v4"
)

// post request to /logs

func injestLogs(e *echo.Group) {
	e.POST("", controllers.HandleInjestLogs)
}
