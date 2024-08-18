package blog

import (
	render "go-cms/utils/render"

	"github.com/gin-gonic/gin"
)

//	func render(c *gin.Context, templateName string, templateDir string, layoutPath string, layoutName string, data gin.H) {
//		templates := template.Must(template.ParseFiles(layoutPath, filepath.Join(templateDir, templateName)))
//		if err := templates.ExecuteTemplate(c.Writer, layoutName, data); err != nil {
//			c.AbortWithError(500, err)
//		}
//	}
func RegisterRoutes(r *gin.Engine, appName string, theme string, cmsName string, cmsVersion string, templateDir string, layoutPath string, layoutName string) {
	// Load the blog templates
	// r.LoadHTMLGlob("./modules/blog/templates/*")

	// Define a route for the blog page
	// r.GET("/blog", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "blog.html", gin.H{
	// 		"Title": "Blog Page",
	// 	})
	// })

	r.GET("/blog", func(c *gin.Context) {

		title := appName + " | Blog"
		render.Render(c,
			"blog.html",
			templateDir,
			layoutPath,
			layoutName,
			gin.H{
				"AppName":    appName,
				"Title":      title,
				"Content":    "Blog Page",
				"Theme":      theme,
				"CmsName":    cmsName,
				"CmsVersion": cmsVersion,
			})
	})
}
