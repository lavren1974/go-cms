package htmx

import (
	render "github.com/lavren1974/go-cms/utils/render"
	structs "github.com/lavren1974/go-cms/utils/structs"

	"html/template"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

const TemplateDir = "./modules/examples/htmx/templates"

var todoList = []string{} // In-memory list to store tasks

func RegisterRoutes(r *gin.Engine, p structs.ModuleParams) {

	templateLayout := structs.TemplateLayout{
		TemplateDir: TemplateDir,
		LayoutPath:  p.LayoutPath,
		LayoutName:  p.LayoutName,
	}

	// Parse all templates in the directory
	r.SetHTMLTemplate(template.Must(template.ParseGlob(filepath.Join(TemplateDir, "*.html"))))

	r.GET("/examples/htmx", func(c *gin.Context) {

		title := p.AppName + " | Htmx"
		render.Render(c,
			"htmx.html",
			templateLayout,
			gin.H{
				"AppName":    p.AppName,
				"Title":      title,
				"Content":    "Htmx Page!",
				"Theme":      p.Theme,
				"CmsName":    p.CmsName,
				"CmsVersion": p.CmsVersion,
			})
	})

	// Route for handling htmx POST requests
	r.POST("/examples/htmx", func(c *gin.Context) {
		// Respond with a simple HTML snippet
		c.HTML(200, "partial.html", gin.H{
			"Message": "This is dynamic content loaded via HTMX!",
		})
	})

	r.GET("/examples/htmx/todo", func(c *gin.Context) {

		title := p.AppName + " | To-Do List"
		render.Render(c,
			"todo.html",
			templateLayout,
			gin.H{
				"AppName":    p.AppName,
				"Title":      title,
				"Theme":      p.Theme,
				"CmsName":    p.CmsName,
				"CmsVersion": p.CmsVersion,
				"TodoList":   todoList,
			})
	})

	// Add a new to-do item via HTMX POST

	r.POST("/examples/htmx/todo/add", func(c *gin.Context) {
		task := c.PostForm("task")
		if task != "" {
			todoList = append(todoList, task)
			index := len(todoList) - 1 // Get the index of the new task
			c.HTML(http.StatusOK, "item.html", gin.H{
				"Task":  task,
				"Index": index,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Task cannot be empty"})
		}
	})

	// Clear all to-do items
	r.POST("/examples/htmx/todo/clear", func(c *gin.Context) {
		todoList = []string{}
		c.Status(200)
	})

	r.POST("/examples/htmx/todo/delete", func(c *gin.Context) {
		index := c.PostForm("index") // Get the index of the item to delete
		idx, err := strconv.Atoi(index)
		if err != nil || idx < 0 || idx >= len(todoList) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid index"})
			return
		}

		// Remove the item from the list
		todoList = append(todoList[:idx], todoList[idx+1:]...)

		c.Status(http.StatusOK) // Return a successful response
	})
}
