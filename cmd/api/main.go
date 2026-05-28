package main

import (
	"log"

	"github.com/Zvoook/lifeforge/internal/app"
)

func main() {
	application := app.New()

	if err := application.Run(); err != nil {
		log.Fatalf("Ошибка при запуске приложения: %v", err)
	}
}