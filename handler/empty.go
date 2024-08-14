package handler

import (
	"rosterize/view/components"

	"github.com/labstack/echo/v4"
)

type EmptyHandler struct {}

func (h EmptyHandler) GetEmpty (c echo.Context) error {
    return render(c, components.Empty())
}
