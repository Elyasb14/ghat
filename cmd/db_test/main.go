package main

import (
	"database/sql"
	"fmt"

	_ "github.com/tursodatabase/go-libsql"
)

func run() (err error) {
	db, err := sql.Open("libsql", "file:///tmp/test.db")
	if err != nil {
		return err
	}

	defer db.Close()

	db.Exec(`CREATE TABLE IF NOT EXISTS chats (
		time TIMESTAMP NOT NULL,
		message TEXT NOT NULL,
		ip_address TEXT NOT NULL
	)`)

	_, err = db.Exec("INSERT INTO chats (time, message, ip_address) VALUES (?, ?, ?)", 103201, "hello second", "10.0.1.1") 
	if err != nil {
		return err
	}

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
