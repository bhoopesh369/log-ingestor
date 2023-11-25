package controllers

import (
	"fmt"

	"github.com/bhoopesh369/log-injestor/services"
	"github.com/labstack/echo/v4"
)

func HandleInjestLogs(c echo.Context) error {
	fmt.Println("log")

	go services.ProducerService(c)

	services.ConsumerService(c)

	// return c.String(200, "ok")
	return nil
}
