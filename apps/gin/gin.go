package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	config "go-cms/utils/config"
	render "go-cms/utils/render"

	"github.com/gin-gonic/gin"

	"go-cms/modules/blog"
	"go-cms/modules/book"
)

func main() {

	// Получаем имя приложения, для загрузки локального конфигурационного файла
	executableName := filepath.Base(filepath.Clean(os.Args[0]))

	//fmt.Println("Application Name:", executableName)

	// Убираем расширение из имени
	if runtime.GOOS == "windows" {
		executableName = strings.TrimSuffix(executableName, ".exe")
	}

	//fmt.Println("Application Name:", executableName)

	// Load global config
	globalConfig, err := config.LoadConfigGlobal("./global.toml")
	if err != nil {
		log.Fatalf("Error loading global config: %v", err)
	}
	fmt.Printf("Global Config: %+v\n", globalConfig)
	log.Println(globalConfig.App.Name)
	log.Println(globalConfig.App.Version)

	pathLocalConfig := "./apps/" + executableName + "/config.toml"

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

	// Load templates from multiple directories
	tmpl := template.Must(template.ParseGlob("./apps/gin/views/*.html"))
	tmpl = template.Must(tmpl.ParseGlob("./modules/book/templates/*.html"))
	tmpl = template.Must(tmpl.ParseGlob("./modules/blog/templates/*.html"))

	// Load HTML templates, including the base layout
	//  r.SetHTMLTemplate(template.Must(template.ParseGlob("layouts/*.html")).
	//  AddParseTree("book.html", template.Must(template.New("book.html").ParseFiles(filepath.Join("book/templates", "book.html")))).
	//  AddParseTree("blog.html", template.Must(template.New("blog.html").ParseFiles(filepath.Join("blog/templates", "blog.html")))))

	// Set the templates in the Gin engine
	r.SetHTMLTemplate(tmpl)

	// Serve static files from the 'static' directory
	r.Static("/static", "./static")

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
			"./apps/gin/views",
			"./apps/gin/views/layout.html",
			"layout.html",
			gin.H{
				"AppName": "Gin Bootstrap App!!!!",
			})
	})

	// Register routes from the blog module
	blog.RegisterRoutes(r)

	// Register routes from the book module
	book.RegisterRoutes(r)

	// Start the web server on port 8080
	r.Run(":8080")

}
