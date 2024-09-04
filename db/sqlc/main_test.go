package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jpmoraess/moraes-bank/util"
	"log"
	"os"
	"testing"
)

var testQueries *Queries
var pool *pgxpool.Pool

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	dbConfig, err := pgxpool.ParseConfig(config.DBSource)
	if err != nil {
		log.Fatal(err)
	}

	pool, err = pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	testQueries = New(pool)

	os.Exit(m.Run())
}
