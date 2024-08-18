package todos

import (
	render "go-cms/utils/render"
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

func RegisterRoutes(r *gin.Engine, appName string, theme string, cmsName string, cmsVersion string, templateDir string, layoutPath string, layoutName string) {
	// // Load the book templates
	// r.LoadHTMLGlob("./modules/book/templates/*")

	// Define a route for the book page
	// r.GET("/book", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "book.html", gin.H{
	// 		"Title": "Book Page",
	// 	})
	// })

	r.GET("/todos", func(c *gin.Context) {

		title := appName + " | Todos"
		render.Render(c,
			"todos.html",
			templateDir,
			layoutPath,
			layoutName,
			gin.H{
				"AppName":    appName,
				"Title":      title,
				"Content":    "Todos Page",
				"todos":      todos,
				"Theme":      theme,
				"CmsName":    cmsName,
				"CmsVersion": cmsVersion,
			})
	})

	r.GET("/todos/:id/edit", func(c *gin.Context) {

		title := appName + " | Todos"

		idParam := c.Param("id")
		id, _ := strconv.Atoi(idParam)
		// if err != nil {
		// 	render.Render(c,
		// 		"todos.html",
		// 		templateDir,
		// 		layoutPath,
		// 		layoutName,
		// 		gin.H{
		// 			"AppName":    appName,
		// 			"Title":      title,
		// 			"Content":    "Todos Page",
		// 			"todos":      todos,
		// 			"Theme":      theme,
		// 			"CmsName":    cmsName,
		// 			"CmsVersion": cmsVersion,
		// 		})
		// 	return
		// }
		for _, todo := range todos {
			if todo.ID == id {
				// c.HTML(http.StatusOK, "edit.html", gin.H{
				// 	"todo": todo,
				// })
				render.Render(c,
					"edit.html",
					templateDir,
					layoutPath,
					layoutName,
					gin.H{
						"AppName":    appName,
						"Title":      title,
						"Content":    "Edit Page",
						"todo":       todo,
						"Theme":      theme,
						"CmsName":    cmsName,
						"CmsVersion": cmsVersion,
					})
				return
			}
		}

	})

	// // Show the edit page
	// r.GET("/todos/:id/edit", func(c *gin.Context) {
	// 	idParam := c.Param("id")
	// 	id, err := strconv.Atoi(idParam)
	// 	if err != nil {
	// 		c.HTML(http.StatusBadRequest, "todos.html", gin.H{
	// 			"error": "Invalid ID",
	// 			"todos": todos,
	// 		})
	// 		return
	// 	}

	// 	for _, todo := range todos {
	// 		if todo.ID == id {
	// 			c.HTML(http.StatusOK, "edit.html", gin.H{
	// 				"todo": todo,
	// 			})
	// 			return
	// 		}
	// 	}

	// 	c.HTML(http.StatusNotFound, "todos.html", gin.H{
	// 		"error": "Todo not found",
	// 		"todos": todos,
	// 	})
	// })

	// // Display the list of todos
	// r.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.html", gin.H{
	// 		"todos": todos,
	// 	})
	// })

	// Create a new todo
	r.POST("/todos", func(c *gin.Context) {
		title := c.PostForm("title")
		if title == "" {
			c.HTML(http.StatusBadRequest, "todos.html", gin.H{
				"error": "Title cannot be empty",
				"todos": todos,
			})
			return
		}

		todo := Todo{
			ID:        idCounter,
			Title:     title,
			Completed: false,
		}
		idCounter++
		todos = append(todos, todo)

		c.Redirect(http.StatusFound, "/todos")
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
