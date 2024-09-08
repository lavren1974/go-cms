package blog

import (
	render "go-cms/utils/render"
	structs "go-cms/utils/structs"

	"github.com/gin-gonic/gin"
)

//	func render(c *gin.Context, templateName string, templateDir string, layoutPath string, layoutName string, data gin.H) {
//		templates := template.Must(template.ParseFiles(layoutPath, filepath.Join(templateDir, templateName)))
//		if err := templates.ExecuteTemplate(c.Writer, layoutName, data); err != nil {
//			c.AbortWithError(500, err)
//		}
//	}

const TemplateDir = "./modules/examples/blog/templates"

func RegisterRoutes(r *gin.Engine, p structs.ModuleParams) {
	// Load the blog templates
	// r.LoadHTMLGlob("./modules/blog/templates/*")

	// Define a route for the blog page
	// r.GET("/blog", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "blog.html", gin.H{
	// 		"Title": "Blog Page",
	// 	})
	// })

	templateLayout := structs.TemplateLayout{
		TemplateDir: TemplateDir,
		LayoutPath:  p.LayoutPath,
		LayoutName:  p.LayoutName,
	}

	r.GET("/blog", func(c *gin.Context) {

		title := p.AppName + " | Blog"
		render.Render(c,
			"blog.html",
			templateLayout,
			gin.H{
				"AppName":    p.AppName,
				"Title":      title,
				"Content":    "Blog Page",
				"Theme":      p.Theme,
				"CmsName":    p.CmsName,
				"CmsVersion": p.CmsVersion,
			})
	})
}
