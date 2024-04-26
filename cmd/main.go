package main

import (
	"rosterize/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	homeHandler := handler.HomeHandler{}
	app.GET("/", homeHandler.HandlerHomeShow).Name = "home"

	rosterHandler := handler.RosterHandler{}
	app.GET("/roster", rosterHandler.GetRoster).Name = "roster"

	newMemberHandler := handler.NewPlayerHandler{}
	app.POST("/add-player", newMemberHandler.GetNewPlayer).Name = "add-player"

    getTeamCsvHandler := handler.GetTeamCsvHandler{}
    app.POST("/get-team-csv", getTeamCsvHandler.GetCsv).Name = "get-team-csv"

	app.Static("/static", "public").Name = "static"

	app.Start(":8080")
}
