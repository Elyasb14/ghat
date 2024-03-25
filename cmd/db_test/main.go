package main

import (
	"database/sql"
	"fmt"
	"time"
	_ "github.com/tursodatabase/go-libsql"
)

func run() (err error) {
	db, err := sql.Open("libsql", "file:///tmp/test.db")
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

	// for i := 0; i < 10; i++ {
	// 	_, err = db.Exec(fmt.Sprintf("INSERT INTO chats (id, name) VALUES (%d, 'test-%d')", i, i))
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	_, err = db.Exec(fmt.Sprintf("INSERT INTO chats (time, message, ip_address) VALUES (%s, %s, %s)", time.Now(), "hello, world",  "10.0.0.2"))

	rows, err := db.Query("SELECT * FROM chats")
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

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
