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
	"github.com/gofiber/template/html/v2"
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

	// Initialize template engine
	engine := html.New("./apps/"+executableName+"/views", ".html")

	// Create a new Fiber app with the template engine
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/static", "./static")

	// Route to render the index page
	app.Get("/", func(c *fiber.Ctx) error {
		appName := localConfig.App.Name
		return c.Render("index", fiber.Map{
			"AppName": appName,
			"Theme":   localConfig.App.Theme,
		})
	})

	// Start server on port 3000
	//log.Fatal(app.Listen(":3001"))
	log.Fatal(app.Listen(localConfig.App.Port))

}
