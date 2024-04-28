package main

import (
	"rosterize/handler"
	"rosterize/middleware"

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

	createTeamHandler := handler.CreateTeamHandler{}
	app.POST("/submit-team-name", createTeamHandler.CreateTeam).Name = "create-team"

    downloadHandler := handler.DownloadFileHandler{}
    app.GET("/download/:fileName", middleware.DeleteAfterDownload(downloadHandler.DownloadFile)).Name = "download-csv"

	app.Static("/static", "./public").Name = "static"

	app.Start(":8080")
}
