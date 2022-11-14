package db

import (
	"context"
	"log"
)

func InitTables(c context.Context) {
	connection := GetDbConnection()
	_, err := connection.Exec(c, `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			email VARCHAR(255) NOT NULL UNIQUE,
			password VARCHAR(255) NOT NULL
		);
	`)
	if err != nil {
		log.Panic("Error creating tables", err)
	}
}
