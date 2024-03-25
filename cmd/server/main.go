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
    return
	}
}
