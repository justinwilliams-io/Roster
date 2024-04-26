package middleware

import (
	"os"

	"github.com/labstack/echo/v4"
)

func DeleteAfterDownload(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func(c echo.Context) {
			os.Remove(c.Param("id") + ".csv")
		}(c)

        return next(c)
	}
}
