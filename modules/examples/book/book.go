package book

import (
	render "go-cms/utils/render"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, appName string, theme string, cmsName string, cmsVersion string, templateDir string, layoutPath string, layoutName string) {
	// // Load the book templates
	// r.LoadHTMLGlob("./modules/book/templates/*")

	// Define a route for the book page
	// r.GET("/book", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "book.html", gin.H{
	// 		"Title": "Book Page",
	// 	})
	// })

	r.GET("/book", func(c *gin.Context) {

		title := appName + " | Book"
		render.Render(c,
			"book.html",
			templateDir,
			layoutPath,
			layoutName,
			gin.H{
				"AppName":    appName,
				"Title":      title,
				"Content":    "Book Page",
				"Theme":      theme,
				"CmsName":    cmsName,
				"CmsVersion": cmsVersion,
			})
	})
}
