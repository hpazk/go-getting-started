package controllers

import (
	"net/http"

	"github.com/hpazk/go-echo-rest-api/app/helpers"
	"github.com/labstack/echo/v4"
)

type (
	PublicController struct {
	}
)

func (controller PublicController) Routes() []helpers.Route {
	return []helpers.Route{
		{
			Method:  echo.GET,
			Path:    "/public",
			Handler: controller.GetPublic,
		},
	}
}

func (controller PublicController) GetPublic(c echo.Context) error {
	return c.String(http.StatusOK, "tes")
}
