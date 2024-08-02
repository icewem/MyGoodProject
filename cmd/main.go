package main

import (
	constDB "MyGoodProject/internal/constants"
	"MyGoodProject/internal/db"
	handler "MyGoodProject/internal/handlers"
	"database/sql"
	migrate "github.com/rubenv/sql-migrate"
	"log"

	"fmt"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {

	// Подключение к БД
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", constDB.User, constDB.Password, constDB.Dbname, constDB.SSLMode)
	database, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer database.Close() // Закрытие соединения в конце

	err = database.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Успешное подключение к базе данных!")

	// Инициализация базы данных
	db.Init(database)

	migrations := migrate.FileMigrationSource{
		Dir: "internal/migrations", //  путь к папке с миграциями
	}

	// Применяем миграции
	_, err = migrate.Exec(database, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Миграции успешно применены!")

	// Мои маршруты
	http.HandleFunc("/get", handler.GetHandler)
	http.HandleFunc("/post", handler.PostHandler)

	// Запуск сервера
	fmt.Println("Сервер запущен на порту 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Ошибка запуска сервера: %s\n", err)
	}
}
