package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/rodolfochicone/go-finance/api"
	db "github.com/rodolfochicone/go-finance/db/sqlc"
	"log"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:postgres@localhost:5432/go_finance?sslmode=disable"
	port     = ":8000"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(port)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
