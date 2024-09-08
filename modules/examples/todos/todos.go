package todos

import (
	render "go-cms/utils/render"
	structs "go-cms/utils/structs"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Todo represents a task with a title and completion status
type Todo struct {
	ID        int
	Title     string
	Completed bool
}

// In-memory storage for todos
var todos = []Todo{}
var idCounter = 1

func RegisterRoutes(r *gin.Engine, p structs.ModuleParams) {
	templateLayout := structs.TemplateLayout{
		TemplateDir: "./modules/examples/todos/templates",
		LayoutPath:  p.LayoutPath,
		LayoutName:  p.LayoutName,
	}

	r.GET("/todos", func(c *gin.Context) {

		title := p.AppName + " | Todos"
		render.Render(c,
			"todos.html",
			templateLayout,
			gin.H{
				"AppName":    p.AppName,
				"Title":      title,
				"Content":    "Todos Page!!!!!",
				"todos":      todos,
				"Theme":      p.Theme,
				"CmsName":    p.CmsName,
				"CmsVersion": p.CmsVersion,
			})
	})

	r.GET("/todos/:id/edit", func(c *gin.Context) {

		title := p.AppName + " | Todos | Edit"

		idParam := c.Param("id")
		id, _ := strconv.Atoi(idParam)

		for _, todo := range todos {
			if todo.ID == id {
				// c.HTML(http.StatusOK, "edit.html", gin.H{
				// 	"todo": todo,
				// })

				render.Render(c,
					"edit.html",
					templateLayout,
					gin.H{
						"AppName":    p.AppName,
						"Title":      title,
						"Content":    "Edit Page",
						"todo":       todo,
						"Theme":      p.Theme,
						"CmsName":    p.CmsName,
						"CmsVersion": p.CmsVersion,
					})

				return
			}
		}

	})

	// Create a new todo
	r.POST("/todos", func(c *gin.Context) {
		title := c.PostForm("title")
		if title == "" {
			// c.HTML(http.StatusBadRequest, "todos.html", gin.H{
			// 	"error": "Title cannot be empty",
			// 	"todos": todos,
			// })
			// return
			log.Println("post error")

			c.Redirect(http.StatusFound, "/todos")
		} else {

			todo := Todo{
				ID:        idCounter,
				Title:     title,
				Completed: false,
			}

			log.Println(todo)
			idCounter++
			todos = append(todos, todo)

			c.Redirect(http.StatusFound, "/todos")
		}

	})

	// Mark a todo as completed
	r.POST("/todos/:id/complete", func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.HTML(http.StatusBadRequest, "todos.html", gin.H{
				"error": "Invalid ID",
				"todos": todos,
			})
			return
		}

		for i, todo := range todos {
			if todo.ID == id {
				todos[i].Completed = true
				break
			}
		}

		c.Redirect(http.StatusFound, "/todos")
	})

	// Delete a todo
	r.POST("/todos/:id/delete", func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.HTML(http.StatusBadRequest, "todos.html", gin.H{
				"error": "Invalid ID",
				"todos": todos,
			})
			return
		}

		for i, todo := range todos {
			if todo.ID == id {
				todos = append(todos[:i], todos[i+1:]...)
				break
			}
		}

		c.Redirect(http.StatusFound, "/todos")
	})

	// Update an existing todo
	r.POST("/todos/:id/edit", func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.HTML(http.StatusBadRequest, "todos.html", gin.H{
				"error": "Invalid ID",
				"todos": todos,
			})
			return
		}

		title := c.PostForm("title")
		if title == "" {
			c.HTML(http.StatusBadRequest, "edit.html", gin.H{
				"error": "Title cannot be empty",
				"todo":  Todo{ID: id, Title: title},
			})
			return
		}

		for i, todo := range todos {
			if todo.ID == id {
				todos[i].Title = title
				c.Redirect(http.StatusFound, "/todos")
				return
			}
		}

		c.HTML(http.StatusNotFound, "todos.html", gin.H{
			"error": "Todo not found",
			"todos": todos,
		})
	})
}
