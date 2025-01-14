package main

import (
	"rosterize/handler"
	"rosterize/middleware"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()

	app.Use(
		echoMiddleware.Recover(),
        echoMiddleware.Gzip(),
	)

    rosterHandler := handler.RosterHandler{}
	app.GET("/", rosterHandler.GetRoster).Name = "home"

	emptyHandler := handler.EmptyHandler{}
	app.GET("/empty", emptyHandler.GetEmpty).Name = "empty"

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
