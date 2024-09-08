package utils

import (
	"html/template"
	"path/filepath"

	structs "go-cms/utils/structs"

	"github.com/gin-gonic/gin"
)

// func Render2(c *gin.Context, templateName string, templateDir string, layoutPath string, layoutName string, data gin.H) {
// 	templates := template.Must(template.ParseFiles(layoutPath, filepath.Join(templateDir, templateName)))
// 	if err := templates.ExecuteTemplate(c.Writer, layoutName, data); err != nil {
// 		c.AbortWithError(500, err)
// 	}
// }

func Render(c *gin.Context, templateName string, t structs.TemplateLayout, data gin.H) {
	templates := template.Must(template.ParseFiles(t.LayoutPath, filepath.Join(t.TemplateDir, templateName)))
	if err := templates.ExecuteTemplate(c.Writer, t.LayoutName, data); err != nil {
		c.AbortWithError(500, err)
	}
}
