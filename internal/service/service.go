package service

import (
	"strings"
	"unicode"
	
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// Функция автоопределения
func AutoConvert(input string) (string, error) {
    input = strings.TrimSpace(input) // убрать пробелы и переносы строк в начале и в конце
    // проверяем, есть ли кириллические буквы
    for _, r := range input {
        if unicode.In(r, unicode.Cyrillic) {
			// если есть значит это обычный текст, конвертируем в код Морзе
            return morse.ToMorse(input), nil
        }
    }
    if strings.ContainsAny(input, ".-") {
        // значит это код Морзе, конвертируем в текст
        return morse.ToText(input), nil
    }
    return input, nil
}