package middleware

import (
	"rosterize/util"

	"github.com/labstack/echo/v4"
)

func DeleteAfterDownload(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func(c echo.Context) {
            status := c.Response().Status

            if status == 200 {
                util.DeleteFile(c.Param("fileName"))
            }
		}(c)

        return next(c)
	}
}
