package db

import (
	"context"
	"log"
)

func InitTables(c context.Context) {
	connection := GetDbConnection()
	_, errUsers := connection.Exec(c, `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			email VARCHAR(255) NOT NULL UNIQUE,
			password VARCHAR(255) NOT NULL
		);
	`)
	if errUsers != nil {
		log.Panic("Error creating users table", errUsers)
	}

	_, errTokens := connection.Exec(c, `
		CREATE TABLE IF NOT EXISTS user_tokens (
			id SERIAL PRIMARY KEY,
			token VARCHAR(255) NOT NULL,
			user_id INT NOT NULL,
			FOREIGN KEY (user_id)
				REFERENCES users(id)
		);
	`)

	if errTokens != nil {
		log.Panic("Error creating tokens table", errTokens)
	}
}
