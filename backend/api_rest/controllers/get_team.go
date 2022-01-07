package controllers

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/JuanMillan7818/FUT21/api_rest/models"
	"github.com/JuanMillan7818/FUT21/db"
	"github.com/labstack/echo/v4"
)

func GetPlayersTeam(e echo.Context) (err error) {
	data := new(models.TeamRequest)
	if err := e.Bind(data); err != nil {
		return err
	}
	data.Name = strings.ToLower(data.Name)
	res := db.GetTeamPlayers(data)

	return e.JSON(http.StatusOK, res)
}

func GetPlayersNames(e echo.Context) (err error) {
	search := e.QueryParam("search")
	order := e.QueryParam("order")
	page := e.QueryParam("page")

	if search == "" && order == "" {

	}
	pageIntm, err := strconv.ParseInt(page, 10, 0)
	if err != nil {
		log.Panic(err)
	}

	res := db.GetNamesPlayers(search, order, int(pageIntm))

	return e.JSON(http.StatusOK, res)
}
