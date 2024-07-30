package db

import (
	"database/sql"
	"fmt"
	"log"
)

// DB - глобальная переменная для соединения с базой данных
var DB *sql.DB

// Инициализация базы данных
func Init(database *sql.DB) {
	DB = database
}

// Пример функции для получения данных
func GetData() ([]string, error) {
	var result []string

	rows, err := DB.Query("SELECT id, name, age FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Обработка полученных данных
	for rows.Next() {
		var id int
		var name string
		var age int

		// Сканируем строку результатов в переменные
		if err := rows.Scan(&id, &name, &age); err != nil {
			log.Fatal(err)
		}
		// Форматируем запись и добавляем в срез
		record := fmt.Sprintf("ID: %d, Name: %s, Age: %d", id, name, age)
		result = append(result, record)
	}

	// Проверяем на наличие ошибок после завершения итерации
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return result, nil
}

// Заготовка под добавление данных
func AddData(name string) error {
	query := "INSERT INTO users (name) VALUES ($1)"
	_, err := DB.Exec(query, name)
	if err != nil {
		return err
	}
	return nil
}
