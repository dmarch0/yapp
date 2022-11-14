package boot

import (
	"context"
	"yapp/core/internal/db"
	"yapp/core/internal/http/server"
)

func Start() {
	c := context.Background()

	dbConnection := db.ConnectToDb(c)
	defer dbConnection.Close(c)
	db.InitTables(c)

	server.StartServer()
}
