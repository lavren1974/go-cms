package book

import (
	render "go-cms/utils/render"
	structs "go-cms/utils/structs"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, p structs.ModuleParams) {

	templateLayout := structs.TemplateLayout{
		TemplateDir: "./modules/examples/book/templates",
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
				"Content":    "Book Page",
				"Theme":      p.Theme,
				"CmsName":    p.CmsName,
				"CmsVersion": p.CmsVersion,
			})
	})
}
