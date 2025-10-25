package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	// Создаём логгер
	logger := log.New(os.Stdout, "MORSE_SERVER: ", log.LstdFlags|log.Lshortfile)

	// Создаём сервер с помощью пакета server
	srv := server.NewServer(logger)

	// Логируем старт
	logger.Println("Сервер запущен на http://localhost:8080")

	// Запускаем сервер
	if err := srv.Server.ListenAndServe(); err != nil {
		logger.Fatal("Ошибка при запуске сервера: ", err)
	}
}
