package main

import (
	"github.com/heroku/go-getting-started/app/routes"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	routes.DefineAPIRoutes()
}
