package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	// Создаём логгер
	logger := log.New(os.Stdout, "MORSE_SERVER: ", log.LstdFlags|log.Lshortfile)

	// Создаём сервер
	srv := server.NewServer(logger)

	// Логируем старт
	logger.Println("Сервер запускается на http://localhost:8080")

	// Запускаем сервер
	err := srv.Server.ListenAndServe()
	if err != nil {
		srv.Logger.Fatal("Ошибка при запуске сервера: ", err)
	}
}
