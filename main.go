package main

import (
	"WhatDownloadGo/handlers"
	"WhatDownloadGo/repositories"
	"WhatDownloadGo/services"
)

func main() {
	// Создание репозитория
	repository := &repositories.HTTPLocationRepository{}

	// Создание сервиса с внедрением зависимости репозитория
	service := services.NewLocationService(repository)

	// Создание обработчика с внедрением зависимости сервиса
	handlers := handlers.NewHandlers(service)

	// Запуск обработчика
	handlers.Run()
}
