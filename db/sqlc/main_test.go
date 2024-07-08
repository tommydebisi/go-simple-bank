package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	driver = "postgres"
	source = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var (
	testQueries *Queries
	testDb      *sql.DB
)

func TestMain(m *testing.M) {
	testDb, err := sql.Open(driver, source)
	if err != nil {
		log.Fatal("unable to establish a db connection:", err)
	}

	testQueries = New(testDb)

	os.Exit(m.Run())
}
