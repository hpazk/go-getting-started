package main

import (
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/heroku/go-getting-started/app/helpers"
	"github.com/heroku/go-getting-started/app/routes"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/joho/godotenv"
)

func main() {
	// routes.DefineAPIRoutes()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Define API wrapper
	api := echo.New()
	api.Validator = &helpers.CustomValidator{Validator: validator.New()}
	api.Use(middleware.Logger())
	api.Use(middleware.Recover())
	// CORS middleware for API endpoint.
	api.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	// db := database.GetInstance()
	// err = database.GetMigrations(db)
	// if err != nil {
	// 	fmt.Println("migrations failed.", err)
	// } else {
	// 	fmt.Println("Migrations did run successfully")
	// }

	// m := database.GetMigrations(db)
	// err = m.Migrate()
	// if err == nil {
	// 	print("Migrations did run successfully")
	// } else {
	// 	print("migrations failed.", err)
	// }
	routes.DefineAPIRoutes(api)

	server := echo.New()
	server.Any("/*", func(c echo.Context) (err error) {
		req := c.Request()
		res := c.Response()
		if req.URL.Path[:4] == "/api" {
			api.ServeHTTP(res, req)
		}

		return
	})
	server.Logger.Fatal(server.Start(":" + os.Getenv("PORT")))
	// e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
