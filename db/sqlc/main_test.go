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
var testQueries *Queries

func TestMain(m *testing.M) {
	connection, err := sql.Open(driver, source)
	if err != nil {
		log.Fatal("unable to establish a db connection:", err)
	}

	testQueries = New(connection)

	os.Exit(m.Run())
}