package handler

import (
	"rosterize/model"
	"rosterize/view/roster"
	"strconv"

	"github.com/labstack/echo/v4"
)

type NewPlayerHandler struct{}

func (h NewPlayerHandler) GetNewPlayer(c echo.Context) error {
	playerNumber, err := strconv.Atoi(c.Request().FormValue("player_number"))
	if err != nil {
		playerNumber = 0
	}

	player := model.Player{
		NameOnShirt:  c.Request().FormValue("name_on_shirt"),
		PlayerNumber: playerNumber,
		ItemType:     c.Request().FormValue("item_type"),
		Size:         c.Request().FormValue("size"),
	}

	c.Response().Header().Add("HX-Trigger-After-Swap", "showRoster")

	return render(c, roster.ShowPlayer(player))
}
