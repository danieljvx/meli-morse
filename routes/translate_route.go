package routes

import (
	"github.com/danieljvx/meli-morse/controllers"
	"github.com/gofiber/fiber/v2"
)

func TranslateRoute(router fiber.Router) {
	translate := router.Group("/translate")
	translate.Post("/decodeBits", controllers.DecodeBits)
	translate.Post("/2text", controllers.ToText)
	translate.Post("/2morse", controllers.ToMorse)
}