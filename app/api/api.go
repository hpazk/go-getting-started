package api

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func Api() {
	e := echo.New()
	e.GET("/api", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
