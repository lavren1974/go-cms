package book

import (
	"github.com/gofiber/fiber/v2"
)

func bookIndexHandler(c *fiber.Ctx) error {
	return c.Render("book/book", fiber.Map{
		"Title": "Book List",
		"Books": []string{"Book 1", "Book 2", "Book 3"},
	})
}

func bookDetailHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.Render("book/detail", fiber.Map{
		"Title": "Book Detail",
		"Book":  "Book " + id,
	})
}
