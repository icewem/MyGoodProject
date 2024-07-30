package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// PostHandler Обработчик POST заглушка на будущее
func PostHandler(w http.ResponseWriter, r *http.Request) {
	// Убедитесь, что запрос метод POST
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	// Декодировать JSON тело запроса
	var data map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Ошибка при декодировании JSON", http.StatusBadRequest)
		return
	}

	// Пример обработки данных (выводим в консоль)
	fmt.Printf("Полученные данные: %+v\n", data)

	// Ответ клиенту
	response := map[string]string{"message": "Это POST запрос", "received": fmt.Sprintf("%v", data)}

	// Установите заголовок ответа
	w.Header().Set("Content-Type", "application/json")

	// Верните JSON ответ
	json.NewEncoder(w).Encode(response)
}
