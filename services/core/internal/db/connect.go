package db

import (
	"context"
	"log"
	"yapp/core/configs"

	"github.com/jackc/pgx/v5"
)

func ConnectToDb(c *context.Context) *pgx.Conn {
	config := configs.GetConfig()
	conn, err := pgx.Connect(*c, config.PostgreSQLConfig.Uri)
	if err != nil {
		log.Fatal(err)
	}
	pingError := conn.Ping(*c)
	if pingError != nil {
		log.Fatal(err)
	}
	log.Println("Connected to DB")
	return conn
}
