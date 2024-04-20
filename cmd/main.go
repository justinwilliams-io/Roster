package main

import (
	"rosterize/handler"

	"github.com/labstack/echo/v4"
)

func main() {
    app := echo.New()

    homeHandler := handler.HomeHandler{}
    app.GET("/", homeHandler.HandlerHomeShow)

    rosterHandler := handler.RosterHandler{}
    app.GET("/roster", rosterHandler.GetRoster)

    newMemberHandler := handler.NewMemberHandler{}
    app.GET("/new-player", newMemberHandler.GetNewTeamMember)

    app.Static("/static", "public")

    app.Start(":8080")
}
