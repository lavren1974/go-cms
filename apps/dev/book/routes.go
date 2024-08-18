package book

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	bookGroup := app.Group("/book")
	bookGroup.Get("/", bookIndexHandler)
	bookGroup.Get("/:id", bookDetailHandler)
}
