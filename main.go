package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jpmoraess/moraes-bank/api"
	db "github.com/jpmoraess/moraes-bank/db/sqlc"
	"log"
)

const (
	connString    = "postgres://postgres:postgres@localhost:5432/postgres"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Fatal(err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}

	store := db.NewStore(pool)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
