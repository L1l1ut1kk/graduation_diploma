package models

import (
	"database/sql"
	"log"
)

func CreateDatabase() {
	var db *sql.DB

	// Указываем параметры подключения к базе данных
	var conninfo string = "user=postgres password=postgres dbname=postgres sslmode=disable"
	conn, err := sql.Open("postgres", conninfo)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	dbName := "diplom_rob"
	db = conn

	// Проверяем существование базы данных
	var exists bool
	err = conn.QueryRow("SELECT EXISTS (SELECT FROM pg_database WHERE datname = $1)", dbName).Scan(&exists)
	if err != nil {
		log.Fatal(err)
	}

	if !exists {
		// Создаем базу данных, если она еще не существует
		_, err = conn.Exec("CREATE DATABASE " + dbName)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Database created successfully")

	// Подключаемся к базе данных your_database_name
	conninfo = "user=postgres password=postgres dbname=" + dbName + " sslmode=disable"
	conn, err = sql.Open("postgres", conninfo)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Создаем таблицу users, если она еще не существует
	_, err = conn.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username TEXT,	
			password TEXT,
			name TEXT,
			surname TEXT,
			mail TEXT
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Создаем таблицу cards, если она еще не существует
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS cards (
			id SERIAL PRIMARY KEY,
			name TEXT
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Создаем таблицу sequences, если она еще не существует
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS sequences (
			id SERIAL PRIMARY KEY,
			user_id INT,
			card_1_id INT,
			card_2_id INT,
			created_at INT
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Tables created successfully")
}

// func PrintUsers(c *gin.Context) {
// 	var db *sql.DB

// 	// Указываем параметры подключения к базе данных
// 	var conninfo string = "user=postgres password=postgres dbname=diplom_rob sslmode=disable"
// 	conn, err := sql.Open("postgres", conninfo)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer conn.Close()

// 	db = conn

// 	// Выполняем запрос на выборку данных из таблицы users
// 	rows, err := db.Query("SELECT * FROM users")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

// 	// Создаем срез для хранения пользователей
// 	users := []User{}

// 	// Обрабатываем результаты запроса
// 	for rows.Next() {
// 		var user User
// 		err = rows.Scan(&user.ID, &user.Username, &user.Password, &user.Name, &user.Surname, &user.Mail)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		users = append(users, user)
// 	}

// 	if err = rows.Err(); err != nil {
// 		log.Fatal(err)
// 	}

// 	// Кодируем пользователей в JSON
// 	jsonData, err := json.Marshal(users)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Выводим JSON
// 	c.JSON(http.StatusOK, jsonData)
// 	fmt.Println(string(jsonData))
// }
