package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// Service — слой бизнес-логики приложения.
type Service struct{}

// New — конструктор для создания нового экземпляра сервиса.
func New() *Service {
	return &Service{}
}

// EncodeToMorse — кодирует текст в азбуку Морзе.
func (s *Service) EncodeToMorse(text string) string {
	return morse.ToMorse(text)
}

// DecodeFromMorse — декодирует азбуку Морзе в обычный текст.
func (s *Service) DecodeFromMorse(code string) string {
	return morse.ToText(code)
}

// Convert — определяет тип входной строки и возвращает результат преобразования.
// Если передан обычный текст → возвращает код Морзе,
// если передан код Морзе → возвращает текст.
func (s *Service) Convert(input string) (string, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return "", errors.New("пустая строка")
	}

	if isMorseCode(input) {
		result := s.DecodeFromMorse(input)
		if result == "" {
			return "", errors.New("ошибка при декодировании из Морзе")
		}
		return result, nil
	}

	result := s.EncodeToMorse(input)
	if result == "" {
		return "", errors.New("ошибка при кодировании в Морзе")
	}
	return result, nil
}

// isMorseCode — определяет, является ли строка кодом Морзе.
// Возвращает true, если строка состоит только из '.', '-', ' ', '/'.
func isMorseCode(s string) bool {
	for _, r := range s {
		switch r {
		case '.', '-', ' ', '/', '\n', '\r', '\t':
			continue
		default:
			return false
		}
	}
	return true
}
