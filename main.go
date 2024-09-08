package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jpmoraess/moraes-bank/api"
	db "github.com/jpmoraess/moraes-bank/db/sqlc"
	"github.com/jpmoraess/moraes-bank/util"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	dbConfig, err := pgxpool.ParseConfig(config.DBSource)
	if err != nil {
		log.Fatal(err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	store := db.NewStore(pool)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
