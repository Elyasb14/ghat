package database

import (
	"database/sql"
	"time"
	// "log"

	_ "github.com/tursodatabase/go-libsql"
	//server "ebianchi/ghat/pkg/server"
)

var Db *sql.DB

// url can be "file:///tmp/ghat.db" or an actual url
// use the env var DB_URL
func InitDB(url string) error {

	db, err := sql.Open("libsql", url)

	if err != nil {
		return err
	}

  defer db.Close()

	db.Exec(`CREATE TABLE IF NOT EXISTS chats (
           id SERIAL PRIMARY KEY,
           time TIMESTAMP NOT NULL,
           message TEXT NOT NULL,
           ip_address TEXT NOT NULL
  )`)

	Db = db
	return nil
}

// TODO: change this to accept *server.Packet
func SaveChat() error {
	stmt, err := Db.Prepare("INSERT INTO chats (time, message, ip_address) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	stmt.Exec(time.Now(), "Hello, world!", "192.168.1.100")
	return nil
}


