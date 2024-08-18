package book

import (
	render "go-cms/utils/render"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// // Load the book templates
	// r.LoadHTMLGlob("./modules/book/templates/*")

	// Define a route for the book page
	// r.GET("/book", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "book.html", gin.H{
	// 		"Title": "Book Page",
	// 	})
	// })

	r.GET("/book", func(c *gin.Context) {
		render.Render(c,
			"book.html",
			"./modules/book/templates",
			"./apps/gin/views/layout.html",
			"layout.html",
			gin.H{
				"Title": "Book Page",
			})
	})
}
