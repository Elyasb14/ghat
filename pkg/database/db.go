package database

import (
  "database/sql"

  _ "github.com/tursodatabase/go-libsql"
  //_ "modernc.org/sqlite"
)

var Db *sql.DB

func InitDB(url string) error {
  
}
