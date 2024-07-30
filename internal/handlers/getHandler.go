package handlers

import (
	"MyGoodProject/internal/db"
	"encoding/json"
	"net/http"
)

// GetHandler Обработчик GET
func GetHandler(w http.ResponseWriter, r *http.Request) {
	// Проверяем на GET
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	name, err := db.GetData()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(name)
}
