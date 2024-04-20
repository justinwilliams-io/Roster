package handler

import (
	"rosterize/model"
	"rosterize/view/roster"

	"github.com/labstack/echo/v4"
)

type RosterHandler struct {}

func(h RosterHandler) GetRoster(c echo.Context) error {
    team := model.Team{
        Name: "Team 1",
    }
    return render(c, roster.Show(team))
}
