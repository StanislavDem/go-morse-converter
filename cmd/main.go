package main

import (
    "log"
	"os"

    "github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	// логгер, вывод в консоль INFO, дата, время 
    logger := log.New(os.Stdout, "INFO: ", log.LstdFlags)
	// создаем сервер из пакета server (func NewServer)
    srv := server.NewServer(logger)
	// вывод о запуске
    logger.Println("Сервер запущен на порту 8080")
	// если ошибка при запуске, то вывод ошибки и завершение программы
    err := srv.HTTP.ListenAndServe()
    if err != nil {
        logger.Fatal(err)
    }
}