package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

//1. Создаем структуру сервера с полями для логгера (log.Logger) и http-сервера (http.Server).

type Server struct {
	Logger *log.Logger
	server *http.Server
}

//2. Создаем функцию, в которой нужно создать http-роутер.
//   Функция принимает log.Logger и возвращает экземпляр структуры вашего сервера.

func NewServer(logger *log.Logger) *Server {
	//2.1. Создаем http-роутер
	mux := http.NewServeMux()

	//2.2. Регистрируем хендлеры
	mux.HandleFunc("/", handlers.Main)         // возвращает HTML из файла index.html.
	mux.HandleFunc("/upload", handlers.Upload) // конвертирует морзе<->текст из html-форму из файла index.html

	//2.3. Создаем экземпляр структуры http.Server.
	serv := &http.Server{
		Addr:         ":8080",          // используем порт 8080
		Handler:      mux,              // передайте ваш http-роутер
		ErrorLog:     logger,           // передайте ваш логгер
		ReadTimeout:  5 * time.Second,  // таймаут для чтения. 5 секунд
		WriteTimeout: 10 * time.Second, // таймаут для записи. 10 секунд
		IdleTimeout:  15 * time.Second, // таймаут ожидания следующего запроса. 15 секунд
	}

	//2.4. Возвращаем ссылку на сервер.
	return &Server{
		Logger: logger,
		server: serv,
	}
}
