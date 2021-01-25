package main

import (
	"github.com/heroku/go-getting-started/app/api"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	// port := os.Getenv("PORT")

	// if port == "" {
	// 	log.Fatal("$PORT must be set")
	// }

	// router := gin.New()
	// router.Use(gin.Logger())
	// router.LoadHTMLGlob("templates/*.tmpl.html")
	// router.Static("/static", "static")

	// router.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.tmpl.html", nil)
	// })

	api.Api()

	// router.Run(":" + port)
}
