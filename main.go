package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/nexpictora-pvt-ltd/cnx-backend/api"
	db "github.com/nexpictora-pvt-ltd/cnx-backend/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:6969/ctt_test_001?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start the server:", err)
	}
}
