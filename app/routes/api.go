package routes

import (
	"github.com/heroku/go-getting-started/app/controllers"
	"github.com/heroku/go-getting-started/app/helpers"
	"github.com/labstack/echo/v4"
)

// func DefineAPIRoutes() {
// 	e := echo.New()
// 	e.GET("/", func(c echo.Context) error {
// 		return c.String(http.StatusOK, "Hello, World!")
// 	})
// 	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
// }
func DefineAPIRoutes(e *echo.Echo) {
	controllers := []helpers.Controller{
		// auth.AuthController{},
		// controllers.ProductsController{},
		controllers.PublicController{},
	}

	var routes []helpers.Route
	for _, controller := range controllers {
		routes = append(routes, controller.Routes()...)
	}

	api := e.Group("/api")
	for _, route := range routes {
		switch route.Method {
		case echo.POST:
			{
				api.POST(route.Path, route.Handler, route.Middleware...)
				break
			}
		case echo.GET:
			{
				api.GET(route.Path, route.Handler, route.Middleware...)
				break
			}
		case echo.DELETE:
			{
				api.DELETE(route.Path, route.Handler, route.Middleware...)
				break
			}
		case echo.PUT:
			{
				api.PUT(route.Path, route.Handler, route.Middleware...)
				break
			}
		case echo.PATCH:
			{
				api.PATCH(route.Path, route.Handler, route.Middleware...)
				break
			}
		}
	}
}
