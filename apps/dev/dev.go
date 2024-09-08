package main

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {

	if len(os.Args) == 0 {
		log.Fatal("No arguments provided")
	}

	// Получаем имя приложения, для загрузки локального конфигурационного файла
	appName := getAppName()

	log.Printf("Application Name: %s", appName)

	// if err := Handler(appName); err != nil {
	// 	log.Fatalf("Handler failed: %v", err)
	// }

	Handler(appName)

}

func getAppName() string {
	appName := filepath.Base(filepath.Clean(os.Args[0]))

	// Remove the extension if on Windows
	if runtime.GOOS == "windows" {
		if ext := filepath.Ext(appName); ext == ".exe" {
			appName = strings.TrimSuffix(appName, ext)
		}
	}
	return appName
}
