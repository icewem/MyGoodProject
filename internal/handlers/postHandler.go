package handlers

import (
	"MyGoodProject/internal/db"
	"encoding/json"
	"fmt"
	"net/http"
)

// PostHandler принимает данные и отправляет запрос на запись в БД.
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

	// Извлекаем данные
	likedUser, ok1 := data["liked_user"].(string)
	liker, ok2 := data["liker"].(string)

	// Проверка на успешное приведение типов и пустоту строк
	if !ok1 || !ok2 || len(likedUser) == 0 || len(liker) == 0 {
		http.Error(w, "Некорректные данные", http.StatusBadRequest)
		return
	}

	db.CheckAndAddUsers(likedUser, liker)
	db.AddLike(liker, likedUser)

	//exists, err := db.СheckLikeExists(liker, likedUser)
	//if err != nil {
	//	log.Fatal(err)
	//}

	// Пример обработки данных (консоль)
	fmt.Printf("Имя того, кого лайкаю: %s\n", likedUser)
	fmt.Printf("Имя того, кто получает лайк: %s\n", liker)

	// Отправим подтверждение клиенту
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Лайк успешно отправлен от %s к %s", liker, likedUser)
}
