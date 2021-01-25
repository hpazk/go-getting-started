package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/app/routes"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	// port := os.Getenv("PORT")

	// if port == "" {
	// 	log.Fatal("$PORT must be set")
	// }

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	routes.DefineAPIRoutes()
	// router.Run(":" + port)
}
