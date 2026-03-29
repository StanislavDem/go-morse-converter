package handlers

import (
    "io"
    "net/http"
    "os"
    "path/filepath"
    "time"

    "github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// Корневой эндпоинт ./
func IndexHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "index.html")
}
// Эндпоинт ./upload, принимает файл из формы index.html, конвертирует его содержимое и возвращает результат.
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// парсим форму
    err := r.ParseMultipartForm(10 << 20) // multipart/form-data, лимит 10 MB
    if err != nil {
        http.Error(w, "Ошибка парсинга формы", http.StatusBadRequest)
        return
    }
	// получаем файл
    file, handler, err := r.FormFile("myFile") // из формы index.html
    if err != nil {
        http.Error(w, "Ошибка получения файла", http.StatusInternalServerError)
        return
    }
    defer file.Close()
	// читаем содержимое файла и закрываем
    data, err := io.ReadAll(file)
    if err != nil {
        http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
        return
    }
	// передаем содержимое в func AutoConvert
    result, err := service.AutoConvert(string(data))
    if err != nil {
        http.Error(w, "Ошибка конвертации", http.StatusBadRequest)
        return
    }
	// создаём локальный файл с результатом
    filename := time.Now().UTC().Format("20060102_150405") + filepath.Ext(handler.Filename)
    err = os.WriteFile(filename, []byte(result), 0644)
    if err != nil {
        http.Error(w, "Ошибка записи файла", http.StatusInternalServerError)
        return
    }
	// возвращаем результат
    w.Write([]byte(result))
}