package main

import (
	"github.com/danieljvx/meli-morse/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"net/http"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Static("/", "./client/build")

	api := app.Group("/api")
	api.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"status": http.StatusOK,
			"message": "Api meli-morse by @danieljvx",
		})
	})
	routes.TranslateRoute(api)

	app.Listen(":8000")
}
