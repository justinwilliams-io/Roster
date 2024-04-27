package handler

import (
	"rosterize/view/roster"

	"github.com/labstack/echo/v4"
)

type RosterHandler struct {}

func(h RosterHandler) GetRoster(c echo.Context) error {
    needBase := c.Request().Header.Get("HX-Request") != "true"
    return render(c, roster.Show(needBase))
}
