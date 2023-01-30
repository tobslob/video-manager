package main

import (
	"database/sql"
	"log"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/tobslob/video-manager/api"
	db "github.com/tobslob/video-manager/db/sqlc"
	"github.com/tobslob/video-manager/utils"
)

// @title           Video Manager API
// @version         1.0
// @description     This API enable us to manage videos, metadata and their annotations.

// @contact.name   Kazeem Odutola
// @contact.email  odutola_k@yahoo.ca

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BeererToken
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
