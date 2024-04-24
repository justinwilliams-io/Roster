package handler

import (
	"rosterize/model"
	"rosterize/view/roster"
	"strconv"

	"github.com/labstack/echo/v4"
)

type NewPlayerHandler struct{}

func (h NewPlayerHandler) GetNewPlayer(c echo.Context) error {
	jerseyNumber, err := strconv.Atoi(c.Request().FormValue("jersey_number"))
	if err != nil {
		jerseyNumber = 0
	}

	player := model.Player{
		FirstName:    c.Request().FormValue("first_name"),
		LastName:     c.Request().FormValue("last_name"),
		JerseyNumber: jerseyNumber,
		JerseySize:   c.Request().FormValue("jersey_size"),
	}
    
    c.Response().Header().Add("HX-Trigger", "showRoster")

	return render(c, roster.ShowPlayer(player))
}
