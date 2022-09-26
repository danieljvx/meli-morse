package app

import (
	"github.com/danieljvx/meli-morse/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"net/http"
)

func App() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"status": http.StatusOK,
			"message": "Api meli-morse by @danieljvx",
		})
	})
	routes.TranslateRoute(app)
	return app
}

