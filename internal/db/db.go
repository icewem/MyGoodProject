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

// CheckAndAddUsers Функция для проверки существования пользователей и добавления их в таблицу
func CheckAndAddUsers(usernames ...string) error {
	for _, username := range usernames {
		var exists bool
		err := DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username=$1)", username).Scan(&exists)
		if err != nil {
			return err
		}

		if !exists {
			_, err := DB.Exec("INSERT INTO users (username) VALUES ($1)", username)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// IncrementLikesCount Функция для увеличения счетчика лайков
func IncrementLikesCount(username string) error {
	_, err := DB.Exec("UPDATE users SET likes_count = likes_count + 1 WHERE username = $1", username)
	return err
}

// AddLike Функция для добавления лайка в таблицу likes
func AddLike(liker, likedUser string) error {
	_, err := DB.Exec("INSERT INTO likes (liker_id, liked_user_id) VALUES ((SELECT id FROM users WHERE username = $1), (SELECT id FROM users WHERE username = $2))", liker, likedUser)
	return err
}

// Функция для проверки существования лайка
func СheckLikeExists(likerID, likedUserID int) (bool, error) {
	var likeExists bool
	var userExists bool

	err := DB.QueryRow(`
        SELECT EXISTS (
            SELECT 1 FROM likes 
            WHERE liker_id = $1 AND liked_user_id = $2
        ) AS like_exists,
        EXISTS (
            SELECT 1 FROM users 
            WHERE id = $1 AND likes_count > 0
        ) AS user_exists;
    `, likerID, likedUserID).Scan(&likeExists, &userExists)

	if err != nil {
		return false, err
	}

	// Возвращаем значение, которое говорит, существует ли лайк и есть ли у пользователя лайки
	return !likeExists && userExists, nil
}
