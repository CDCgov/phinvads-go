package database

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	"github.com/CDCgov/phinvads-fhir/internal/config"
)

var (
	Connection *sql.DB
)

func CreateDB(cfg *config.Config) (*sql.DB, error) {
	dsn := cfg.Dsn
	flag.Parse()

	db, err := openDB(*dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return db, nil
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
