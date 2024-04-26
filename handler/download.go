package handler

import "github.com/labstack/echo/v4"

type DownloadFileHandler struct{}

func (h DownloadFileHandler)  DownloadFile(c echo.Context) error {
    id := c.Param("id")

    return c.File(id + ".csv")
}
