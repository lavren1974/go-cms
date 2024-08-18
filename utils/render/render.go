package utils

import (
	"html/template"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func Render(c *gin.Context, templateName string, templateDir string, layoutPath string, layoutName string, data gin.H) {
	templates := template.Must(template.ParseFiles(layoutPath, filepath.Join(templateDir, templateName)))
	if err := templates.ExecuteTemplate(c.Writer, layoutName, data); err != nil {
		c.AbortWithError(500, err)
	}
}
