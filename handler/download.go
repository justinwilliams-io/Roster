package handler

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type DownloadFileHandler struct{}

func (h DownloadFileHandler) DownloadFile(c echo.Context) error {
    fileName := c.Param("fileName")
    _, err :=os.Open("./files/" + fileName)
    if err != nil {
        c.Response().Status = http.StatusNotFound
    }

    return c.Attachment("./files/" + fileName, fileName)
}
