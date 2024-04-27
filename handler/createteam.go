package handler

import (
	"rosterize/view/roster"

	"github.com/labstack/echo/v4"
)

type CreateTeamHandler struct {}

func (h CreateTeamHandler) CreateTeam (c echo.Context) error {
    teamName := c.Request().FormValue("team_name")

    return render(c, roster.ShowRosterForm(teamName))
}
