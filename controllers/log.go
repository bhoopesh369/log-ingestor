package controllers

import (
	"fmt"
	"net/http"

	"github.com/bhoopesh369/log-injestor/config"
	"github.com/bhoopesh369/log-injestor/models"
	"github.com/bhoopesh369/log-injestor/utils"
	"github.com/labstack/echo/v4"
)

func HandleInjestLogs(c echo.Context) error {
	fmt.Println("log")

	log := new(models.Log)
	if err := c.Bind(log); err != nil {
		return err
	}

	db := config.GetDB()

	logCollection := db.Collection(models.LogCollectionName())

	_, err := logCollection.InsertOne(c.Request().Context(), log)

	if err != nil {
		return utils.SendResponse(c, http.StatusInternalServerError, "Error while inserting log")
	}

	return utils.SendResponse(c, http.StatusOK, "Log inserted successfully")
}
