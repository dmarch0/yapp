package db

import (
	"context"
	"log"
)

func InitTables(c context.Context) {
	connection := GetDbConnection()

	usersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			email VARCHAR(255) NOT NULL UNIQUE,
			password VARCHAR(255) NOT NULL
		);
	`
	if _, err := connection.Exec(c, usersTable); err != nil {
		log.Panic("Eror creating users table", err)
	}

	tokensTable := `
		CREATE TABLE IF NOT EXISTS user_tokens (
			id SERIAL PRIMARY KEY,
			token VARCHAR(255) NOT NULL,
			user_id INT NOT NULL,
			FOREIGN KEY (user_id)
				REFERENCES users(id)
		);
	`

	if _, err := connection.Exec(c, tokensTable); err != nil {
		log.Panic("Eror creating tokens table", err)
	}

	charactersTable := `
		CREATE TABLE IF NOT EXISTS characters (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			user_id INT NOT NULL,
			FOREIGN KEY (user_id)
				REFERENCES users(id)
		);
	`

	if _, err := connection.Exec(c, charactersTable); err != nil {
		log.Panic("Eror creating characters table", err)
	}
}
