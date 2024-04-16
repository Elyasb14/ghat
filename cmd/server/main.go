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
	fmt.Println(server)
}
