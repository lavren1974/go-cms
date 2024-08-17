package blog

import (
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes registers the blog routes.
func SetupRoutes(app *fiber.App) {
	// Blog routes
	blogGroup := app.Group("/blog")

	blogGroup.Get("/", getAllBlogs)
	blogGroup.Get("/:id", getBlogByID)
}

func getAllBlogs(c *fiber.Ctx) error {
	// Add logic to retrieve all blogs
	return c.SendString("List of all blog posts")
}

func getBlogByID(c *fiber.Ctx) error {
	// Add logic to retrieve a blog by ID
	id := c.Params("id")
	return c.SendString("Blog post with ID: " + id)
}
