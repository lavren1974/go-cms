package book

import (
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes registers the book routes.
func SetupRoutes(app *fiber.App) {
	// Book routes
	bookGroup := app.Group("/book")

	bookGroup.Get("/", getAllBooks)
	bookGroup.Get("/:id", getBookByID)
}

func getAllBooks(c *fiber.Ctx) error {
	// Add logic to retrieve all books
	return c.SendString("List of all books")
}

func getBookByID(c *fiber.Ctx) error {
	// Add logic to retrieve a book by ID
	id := c.Params("id")
	return c.SendString("Book with ID: " + id)
}
