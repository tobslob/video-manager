package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/tobslob/video-manager/api"
	db "github.com/tobslob/video-manager/db/sqlc"
	"github.com/tobslob/video-manager/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load env", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	if err = server.Start(config.ServerAddress); err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
