package controllers

import "github.com/labstack/echo/v4"

func EndPointMain(e *echo.Echo) {
	e.POST("/api/v1/team", GetPlayersTeam)

	e.GET("/api/v1/players", GetPlayersNames)
}
