package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// Хендлер для корневого эндпоинта '/'
// возвращает HTML из файла index.html

func Main(res http.ResponseWriter, req *http.Request) {
	// Провека что запрос методом GET
	if req.Method != http.MethodGet {
		http.Error(res, "Метод не поддерживается(необходим GET-запрос)", http.StatusMethodNotAllowed)
		return
	}

	// Открываем index.html
	file, err := os.Open("index.html")
	if err != nil {
		http.Error(res, "Не удалось открыть index.html", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Устанавливаем тип контента
	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Копируем содержимое файла в ответ
	_, err = io.Copy(res, file)
	if err != nil {
		http.Error(res, "Ошибка при отправке файла", http.StatusInternalServerError)
		return
	}
}

// Хендлер для эндпоинта '/upload'
func Upload(res http.ResponseWriter, req *http.Request) {
	// Провека что запрос методом POST
	if req.Method != http.MethodPost {
		http.Error(res, "Метод не поддерживается(необходим POST-запрос)", http.StatusMethodNotAllowed)
		return
	}

	// Парсим форму (максимум 10 МБ)
	err := req.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(res, "Ошибка при парсинге формы", http.StatusBadRequest)
		return
	}

	// Получаем файл из формы
	file, header, err := req.FormFile("file")
	if err != nil {
		http.Error(res, "Не удалось получить файл", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Читаем данные из файла
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(res, "Ошибка при чтении файла", http.StatusInternalServerError)
		return
	}

	// Создаем сервис для функции Convert
	svc := service.New()

	// Определяем, что это — морзе или текст, и конвертируем
	result, err := svc.Convert(string(data))
	if err != nil {
		http.Error(res, "Ошибка при конвертации: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Создаём имя для нового файла
	newFileName := fmt.Sprintf("%s_converted%s", time.Now().UTC().Format("20060102_150405"), filepath.Ext(header.Filename))

	// Создаём новый файл для записи результата
	newFile, err := os.Create(newFileName)
	if err != nil {
		http.Error(res, "Ошибка при создании файла", http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	// Записываем результат
	_, err = newFile.Write([]byte(result))
	if err != nil {
		http.Error(res, "Ошибка при записи результата", http.StatusInternalServerError)
		return
	}

	// Возвращаем результат пользователю
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(res, "Результат конвертации:\n")
	fmt.Fprintln(res, result)
}
