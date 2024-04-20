package handler

import (
	"rosterize/model"
	"rosterize/view/home"

	"github.com/labstack/echo/v4"
)

type HomeHandler struct {}

func(h HomeHandler) HandlerHomeShow(c echo.Context) error {
    u := model.User{
        Email: "justin@example.com",
    }
    return render(c, home.Show(u))
}
