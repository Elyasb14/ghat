package main

import (
	"flag"
	"fmt"

	"github.com/Elyasb14/ghat/pkg/server"
)

func main() {
	var port uint
	flag.UintVar(&port, "port", 8080, "port to listen on")
	flag.Parse()
	fmt.Printf("server listening on port %d\n", port)

	server, err := server.NewTCPServer(port)
	if err != nil {
		fmt.Println(err)
	}
	for {
		conn, err := server.Listener.Accept()
		if err != nil {
			fmt.Println(err)
		}
		msg := "hello"
		server.Clients = append(server.Clients, conn)
		for _, client := range server.Clients {
			client.Write([]byte(msg))
		}
	}
}
