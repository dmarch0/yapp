package db

import (
	"context"
	"log"
	"yapp/core/configs"

	"github.com/jackc/pgx/v5"
)

var conn *pgx.Conn

func ConnectToDb(c context.Context) *pgx.Conn {
	config := configs.GetConfig()
	_conn, err := pgx.Connect(c, config.PostgreSQLConfig.Uri)
	if err != nil {
		log.Panic(err)
	}
	pingError := _conn.Ping(c)
	if pingError != nil {
		log.Panic(err)
	}
	log.Println("Connected to DB")
	conn = _conn
	return conn
}

func GetDbConnection() *pgx.Conn {
	if conn == nil {
		log.Panic("Can't get DB connection when it's supposed to be")
	}

	return conn
}
