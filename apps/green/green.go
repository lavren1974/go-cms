package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"go-cms/utils"
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
	globalConfig, err := utils.LoadConfig("./global.toml")
	if err != nil {
		log.Fatalf("Error loading global config: %v", err)
	}
	fmt.Printf("Global Config: %+v\n", globalConfig)

	pathLocalConfig := "./apps/" + executableName + "/config.toml"

	// Load local config
	localConfig, err := utils.LoadConfig(pathLocalConfig)
	if err != nil {
		log.Fatalf("Error loading local config: %v", err)
	}
	fmt.Printf("Local Config: %+v\n", localConfig)
	log.Println(localConfig.App.Name)

}