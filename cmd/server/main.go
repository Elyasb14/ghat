package main

import (
	// "ebianchi/ghat/pkg/server"
	db "ebianchi/ghat/pkg/database"
	"fmt"
)

func main() {
	//server.StartTCPServer()
	err := db.InitDB("file:///tmp/ghat.db")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(db.Db)
	db.SaveChat()
	if err != nil {
		fmt.Println(err)
	}
	db.ShowTable()
}
