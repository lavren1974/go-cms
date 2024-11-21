package main

import (
	"html/template"
	"log"

	config "github.com/lavren1974/go-cms/utils/config"
	render "github.com/lavren1974/go-cms/utils/render"
	structs "github.com/lavren1974/go-cms/utils/structs"

	"github.com/gin-gonic/gin"

	"github.com/lavren1974/go-cms/modules/examples/examples"
	"github.com/lavren1974/go-cms/modules/examples/htmx"
)

func loadConfig(appName string) (*config.GlobalConfig, *config.LocalConfig, error) {
	globalConfig, err := config.LoadConfigGlobal("./global.toml")
	if err != nil {
		return nil, nil, err
	}

	pathLocalConfig := "./apps/" + appName + "/config.toml"
	localConfig, err := config.LoadConfigLocal(pathLocalConfig)
	if err != nil {
		return nil, nil, err
	}

	return globalConfig, localConfig, nil
}

func loadTemplates(appName string) (*template.Template, error) {

	log.Println(appName)
	pathViewsHtml := "./apps/" + appName + "/views/*.html"

	tmpl := template.Must(template.ParseGlob(pathViewsHtml))

	// -----------------------------------------MODULES ADD-----------------------------------------

	var moduleDirs []string
	moduleDirs = append(moduleDirs, examples.TemplateDir)
	moduleDirs = append(moduleDirs, htmx.TemplateDir)

	for _, dir := range moduleDirs {
		dir = dir + "/*.html"
		tmpl = template.Must(tmpl.ParseGlob(dir))
	}

	return tmpl, nil
}

func setupRouter(appName string, globalConfig *config.GlobalConfig, localConfig *config.LocalConfig) *gin.Engine {
	r := gin.Default()

	//pathViewsHtml := "./apps/" + localConfig.App.Name + "/views/*.html"
	pathViews := "./apps/" + appName + "/views"
	pathPublic := "./apps/" + appName + "/public"
	pathLayout := "./apps/" + appName + "/views/layout.html"
	nameLayout := "layout.html"

	tmpl, err := loadTemplates(appName)
	if err != nil {
		log.Fatalf("Error loading templates: %v", err)
	}
	r.SetHTMLTemplate(tmpl)

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

	// -----------------------------------------MODULES ADD-----------------------------------------

	examples.RegisterRoutes(r, moduleParams)
	htmx.RegisterRoutes(r, moduleParams)

	return r
}

func Handler(appName string) {

	globalConfig, localConfig, err := loadConfig(appName)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	r := setupRouter(appName, globalConfig, localConfig)
	r.Run(localConfig.App.Port)

}
