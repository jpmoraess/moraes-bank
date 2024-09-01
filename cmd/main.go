package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jpmoraess/moraes-bank/internal/infra/db"
	"log"
	"time"
)

func main() {
	pool, err := pgxpool.NewWithConfig(context.Background(), databaseConfig())
	if err != nil {
		log.Fatal(err)
	}

	q := db.New(pool)

	err = q.CreateAccount(context.Background(), db.CreateAccountParams{
		Owner:    "Andressa",
		Balance:  2000,
		Currency: "USD",
	})
	if err != nil {

	}

	accounts, err := q.GetAccounts(context.Background(), db.GetAccountsParams{
		Limit:  10,
		Offset: 0,
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, a := range accounts {
		fmt.Printf("Account ID: %+v \n", a)
	}
}

func databaseConfig() *pgxpool.Config {
	config, err := pgxpool.ParseConfig("postgres://postgres:postgres@localhost:5432/postgres")
	if err != nil {
		log.Fatal(err)
	}
	config.MaxConns = 4
	config.MinConns = 1
	config.MaxConnLifetime = time.Hour
	config.MaxConnIdleTime = 5 * time.Minute
	config.HealthCheckPeriod = 1 * time.Minute
	config.ConnConfig.ConnectTimeout = 10 * time.Second

	return config
}
