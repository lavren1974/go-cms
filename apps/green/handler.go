package main

import (
	"html/template"
	"log"

	config "go-cms/utils/config"
	render "go-cms/utils/render"
	structs "go-cms/utils/structs"

	"github.com/gin-gonic/gin"

	"go-cms/modules/examples/blog"
	"go-cms/modules/examples/book"
	"go-cms/modules/examples/todos"
)

func Handler(appName string) {

	// Load global config
	globalConfig, err := config.LoadConfigGlobal("./global.toml")
	if err != nil {
		log.Fatalf("Error loading global config: %v", err)
	}
	// fmt.Printf("Global Config: %+v\n", globalConfig)
	// log.Println(globalConfig.Cms.Name)
	// log.Println(globalConfig.Cms.Version)

	pathLocalConfig := "./apps/" + appName + "/config.toml"

	// Load local config
	localConfig, err := config.LoadConfigLocal(pathLocalConfig)
	if err != nil {
		log.Fatalf("Error loading local config: %v", err)
	}
	//fmt.Printf("Local Config: %+v\n", localConfig)
	// log.Println(localConfig.App.Name)
	// log.Println(localConfig.App.Port)

	// Initialize the Gin router
	r := gin.Default()

	// Load HTML files for templates
	//r.LoadHTMLGlob("./apps/gin/templates/*")

	pathViewsHtml := "./apps/" + appName + "/views/*.html"
	pathViews := "./apps/" + appName + "/views"
	pathPublic := "./apps/" + appName + "/public"
	pathLayout := "./apps/" + appName + "/views/layout.html"
	nameLayout := "layout.html"

	// Load templates from multiple directories
	tmpl := template.Must(template.ParseGlob(pathViewsHtml))
	tmpl = template.Must(tmpl.ParseGlob("./modules/examples/book/templates/*.html"))
	tmpl = template.Must(tmpl.ParseGlob("./modules/examples/blog/templates/*.html"))
	tmpl = template.Must(tmpl.ParseGlob("./modules/examples/todos/templates/*.html"))

	// Set the templates in the Gin engine
	r.SetHTMLTemplate(tmpl)

	// Serve static files from the 'static' directory
	r.Static("/static", "./static")
	r.Static("/public", pathPublic)

	templateLayout := structs.TemplateLayout{
		TemplateDir: pathViews,
		LayoutPath:  pathLayout,
		LayoutName:  nameLayout,
	}

	moduleParams := structs.ModuleParams{
		AppName:    localConfig.App.Name,
		Theme:      localConfig.App.Theme,
		CmsName:    globalConfig.Cms.Name,
		CmsVersion: globalConfig.Cms.Version,
		LayoutPath: pathLayout,
		LayoutName: nameLayout,
	}

	r.GET("/", func(c *gin.Context) {
		render.Render(c,
			"index.html",
			templateLayout,
			gin.H{
				"AppName":    moduleParams.AppName,
				"Title":      moduleParams.AppName,
				"Theme":      moduleParams.Theme,
				"CmsName":    moduleParams.CmsName,
				"CmsVersion": moduleParams.CmsVersion,
			})
	})

	// Register routes from the blog module
	blog.RegisterRoutes(r, moduleParams)
	book.RegisterRoutes(r, moduleParams)
	todos.RegisterRoutes(r, moduleParams)

	// Start the web server on port 8080
	r.Run(localConfig.App.Port)

}
