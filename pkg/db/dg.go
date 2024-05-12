package db

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/tursodatabase/go-libsql"
)

type DataBase struct {
	Db  *sql.DB
	Mut sync.Mutex
}

func InitDB(fd string) *DataBase {
	db, err := sql.Open("libsql", fd)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS messages (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        msg TEXT NOT NULL,
        sender TEXT NOT NULL
  )`)

	return &DataBase{
		Db: db,
	}
}
