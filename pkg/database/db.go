package database

import (
  "database/sql"

  _ "github.com/tursodatabase/go-libsql"
)

var Db *sql.DB

// url can be "file:///tmp/ghat.db" or an actual url
// use the env var DB_URL
func InitDB(url string) error {

  db, err := sql.Open("libsql", url)  

  if err != nil {
    return err
  }
  
  db.Exec(`CREATE TABLE IF NOT EXISTS chats (
           id SERIAL PRIMARY KEY,
           time TIMESTAMP NOT NULL,
           message TEXT NOT NULL,
           ip_address TEXT NOT NULL
  )`)

  Db = db
  return nil
}


