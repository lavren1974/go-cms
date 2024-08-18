package main

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {

	// Получаем имя приложения, для загрузки локального конфигурационного файла
	appName := filepath.Base(filepath.Clean(os.Args[0]))

	//fmt.Println("Application Name:", executableName)

	// Убираем расширение из имени
	if runtime.GOOS == "windows" {
		appName = strings.TrimSuffix(appName, ".exe")
	}
	log.Println(appName)

	Handler(appName)

}
