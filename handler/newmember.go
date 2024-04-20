package handler

import (
	"rosterize/view/components"

	"github.com/labstack/echo/v4"
)

type NewMemberHandler struct {}

func(h NewMemberHandler) GetNewTeamMember(c echo.Context) error {
    return render(c, components.Input("text", "member"))
}
