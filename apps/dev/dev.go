package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	config "go-cms/utils/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var todos []Todo
var idCounter int

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

	// Initialize template engine
	engine := html.New("./apps/"+executableName+"/views", ".html")

	// Create a new Fiber app with the template engine
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Use logger middleware.
	app.Use(logger.New())

	app.Static("/static", "./static")

	// Route to render the index page
	app.Get("/", func(c *fiber.Ctx) error {
		appName := localConfig.App.Name
		return c.Render("index", fiber.Map{
			"AppName": appName,
			"Theme":   localConfig.App.Theme,
			"Todos":   todos,
		})
	})

	app.Get("/todos", func(c *fiber.Ctx) error {
		return c.Render("todo-list", fiber.Map{
			"Todos": todos,
		})
	})

	app.Post("/todos", func(c *fiber.Ctx) error {
		title := c.FormValue("title")
		if title == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Title cannot be empty")
		}
		idCounter++
		newTodo := Todo{ID: idCounter, Title: title, Done: false}
		todos = append(todos, newTodo)
		return c.Render("todo-item", newTodo)
	})

	app.Put("/todos/:id/toggle", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
		}

		for i := range todos {
			if todos[i].ID == id {
				todos[i].Done = !todos[i].Done
				return c.Render("todo-item", todos[i])
			}
		}

		return c.Status(fiber.StatusNotFound).SendString("Todo not found")
	})

	app.Delete("/todos/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
		}

		for i := range todos {
			if todos[i].ID == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.SendStatus(fiber.StatusOK)
			}
		}

		return c.Status(fiber.StatusNotFound).SendString("Todo not found")
	})

	app.Put("/todos/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
		}

		title := c.FormValue("title")
		if title == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Title cannot be empty")
		}

		for i := range todos {
			if todos[i].ID == id {
				todos[i].Title = title
				return c.Render("todo-item", todos[i])
			}
		}

		return c.Status(fiber.StatusNotFound).SendString("Todo not found")
	})

	// Start server on port 3000
	//log.Fatal(app.Listen(":3001"))
	log.Fatal(app.Listen(localConfig.App.Port))

}
