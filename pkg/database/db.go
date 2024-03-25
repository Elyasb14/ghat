package database

import (
	"database/sql"
	"time"
	"fmt"
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
func SaveChat(chat string, ip_addr string) error {
	_, err := Db.Exec("INSERT INTO chats (time, message, ip_address) VALUES (?, ?, ?)", time.Now(), chat, ip_addr) 
	if err != nil {
		return err
	}
	return nil
}

func ShowChats() error {
	rows, err := Db.Query("select * FROM chats")
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var time string
		var message string
		var ip_addr string
		err = rows.Scan(&time, &message, &ip_addr)
		if err != nil {
			return err
		}
		fmt.Printf("time %s, message: %s, ip_addr: %s\n", time, message, ip_addr)
	}
	if rows.Err() != nil {
		return rows.Err()
	}
	return nil
}


