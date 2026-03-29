package server

import (
    "log"
	"net/http"
	"time"

    "github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

// логгер и http-сервер
type Server struct {
    Logger *log.Logger
    HTTP   *http.Server
}

func NewServer(logger *log.Logger) *Server {
	// маршрутизатор
    mux := http.NewServeMux()
	// регистрируем хэндлеры
    mux.HandleFunc("/", handlers.IndexHandler)
    mux.HandleFunc("/upload", handlers.UploadHandler)
	// настройки сервера
    srv := &http.Server{
        Addr:         ":8080",			// порт
        Handler:      mux,				// маршрутизатор
		ErrorLog:     logger,         	// логгер
        ReadTimeout:  5 * time.Second,  // таймаут чтения
        WriteTimeout: 10 * time.Second, // таймаут записи
        IdleTimeout:  15 * time.Second, // таймаут ожидания

    }
    return &Server{
		Logger: logger, 
		HTTP: srv,
	}
}