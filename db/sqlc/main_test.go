package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"testing"
)

const (
	dbSource = "postgres://postgres:postgres@localhost:5432/postgres"
	dbDriver = "postgres"
)

var testQueries *Queries
var pool *pgxpool.Pool

func TestMain(m *testing.M) {
	config, err := pgxpool.ParseConfig("postgres://postgres:postgres@localhost:5432/postgres")
	if err != nil {
		log.Fatal(err)
	}
	pool, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}
	testQueries = New(pool)

	os.Exit(m.Run())
}
