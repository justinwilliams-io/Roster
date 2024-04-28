package handler

import (
	"rosterize/view/home"

	"github.com/labstack/echo/v4"
)

type HomeHandler struct{}

func (h HomeHandler) HandlerHomeShow(c echo.Context) error {
	needBase := c.Request().Header.Get("HX-Request") != "true"
    return render(c, home.Show(needBase))
}
