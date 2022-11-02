package boot

import (
	"yapp/core/internal/db"
	"yapp/core/internal/http/server"
)

func Start() {
	db.ConnectToMongo()
	server.StartServer()
}
