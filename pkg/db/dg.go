package db

import (
	"database/sql"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"log"
)

var db *sql.DB

func InitDB() *sql.DB {
	db, err := sql.Open("libsql", "file:/tmp/db.go")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
