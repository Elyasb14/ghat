package main

import (
	"flag"
	"log"

	"github.com/Elyasb14/ghat/pkg/server"
)

func main() {
	var port uint
	flag.UintVar(&port, "port", 8080, "port to listen on")
	flag.Parse()
	log.Printf("server listening on port %d\n", port)

	tcpServer, err := server.NewTCPServer(port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := tcpServer.Listener.Accept()
		
		if err != nil {
			log.Println(err)
			continue
		}
		
		
		log.Printf("client connected @ ip address %s", conn.RemoteAddr().String())
		go server.HandleClient(conn)
	}
}
