package main

import (
	"fmt"
	"html/template"
	"log"

	config "go-cms/utils/config"
	render "go-cms/utils/render"

	"github.com/gin-gonic/gin"

	"go-cms/modules/examples/blog"
	"go-cms/modules/examples/book"
	"go-cms/modules/examples/todos"
)

func Handler(appName string) {

	fmt.Println("Application Name:", appName)

	// Load global config
	globalConfig, err := config.LoadConfigGlobal("./global.toml")
	if err != nil {
		log.Fatalf("Error loading global config: %v", err)
	}
	fmt.Printf("Global Config: %+v\n", globalConfig)
	log.Println(globalConfig.Cms.Name)
	log.Println(globalConfig.Cms.Version)

	pathLocalConfig := "./apps/" + appName + "/config.toml"

	// Load local config
	localConfig, err := config.LoadConfigLocal(pathLocalConfig)
	if err != nil {
		log.Fatalf("Error loading local config: %v", err)
	}
	//fmt.Printf("Local Config: %+v\n", localConfig)
	log.Println(localConfig.App.Name)
	log.Println(localConfig.App.Port)

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

	// Load HTML templates, including the base layout
	//  r.SetHTMLTemplate(template.Must(template.ParseGlob("layouts/*.html")).
	//  AddParseTree("book.html", template.Must(template.New("book.html").ParseFiles(filepath.Join("book/templates", "book.html")))).
	//  AddParseTree("blog.html", template.Must(template.New("blog.html").ParseFiles(filepath.Join("blog/templates", "blog.html")))))

	// Set the templates in the Gin engine
	r.SetHTMLTemplate(tmpl)

	// Serve static files from the 'static' directory
	r.Static("/static", "./static")
	r.Static("/public", pathPublic)

	// Define a route for the index page
	// r.GET("/", func(c *gin.Context) {
	// 	// Pass the application name to the template
	// 	c.HTML(http.StatusOK, "index.html", gin.H{
	// 		"AppName": "Gin Bootstrap App!!!!",
	// 	})
	// })

	r.GET("/", func(c *gin.Context) {
		render.Render(c,
			"index.html",
			pathViews,
			pathLayout,
			nameLayout,
			gin.H{
				"AppName":    localConfig.App.Name,
				"Title":      localConfig.App.Name,
				"Theme":      localConfig.App.Theme,
				"CmsName":    globalConfig.Cms.Name,
				"CmsVersion": globalConfig.Cms.Version,
			})
	})

	// Register routes from the blog module
	blog.RegisterRoutes(r,
		localConfig.App.Name,
		localConfig.App.Theme,
		globalConfig.Cms.Name,
		globalConfig.Cms.Version,
		"./modules/examples/blog/templates",
		pathLayout,
		nameLayout)

	book.RegisterRoutes(r,
		localConfig.App.Name,
		localConfig.App.Theme,
		globalConfig.Cms.Name,
		globalConfig.Cms.Version,
		"./modules/examples/book/templates",
		pathLayout,
		nameLayout)

	todos.RegisterRoutes(r,
		localConfig.App.Name,
		localConfig.App.Theme,
		globalConfig.Cms.Name,
		globalConfig.Cms.Version,
		"./modules/examples/todos/templates",
		pathLayout,
		nameLayout)

	// Start the web server on port 8080
	r.Run(localConfig.App.Port)

}
