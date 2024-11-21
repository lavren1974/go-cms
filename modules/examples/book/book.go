package book

import (
	render "github.com/lavren1974/go-cms/utils/render"
	structs "github.com/lavren1974/go-cms/utils/structs"

	"github.com/gin-gonic/gin"
)

const TemplateDir = "./modules/examples/book/templates"

func RegisterRoutes(r *gin.Engine, p structs.ModuleParams) {

	templateLayout := structs.TemplateLayout{
		TemplateDir: TemplateDir,
		LayoutPath:  p.LayoutPath,
		LayoutName:  p.LayoutName,
	}

	r.GET("/book", func(c *gin.Context) {

		title := p.AppName + " | Book"
		render.Render(c,
			"book.html",
			templateLayout,
			gin.H{
				"AppName":    p.AppName,
				"Title":      title,
				"Content":    "Book Page!!!!!",
				"Theme":      p.Theme,
				"CmsName":    p.CmsName,
				"CmsVersion": p.CmsVersion,
			})
	})

	// Route for handling htmx POST requests
	r.POST("/book", func(c *gin.Context) {
		// Respond with a simple HTML snippet
		c.HTML(200, "partial.html", gin.H{
			"Message": "This is dynamic content loaded via HTMX!",
		})
	})
}
