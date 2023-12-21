package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var instance *sql.DB

func Conn() error {
	db, err := sql.Open("sqlite3", "./dev.db")
	if err != nil {
		log.Fatal(err)
	}

	instance = db

	return nil
}

func UseQuery() *Queries {
	return New(instance)
}

func Close() error {
	instance.Close()

	return nil
}
