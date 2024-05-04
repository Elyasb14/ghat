package main

import (
	"flag"
	"log"

	"github.com/Elyasb14/ghat/pkg/server"
)

func main() {
	var port, bufSize, maxCons uint
	flag.UintVar(&port, "port", 8080, "port to listen on")
	flag.UintVar(&bufSize, "bufsize", 64, "maximum message buffer size")
	flag.UintVar(&maxCons, "maxcons", 20, "maximum number of allowed connections to the server")
	flag.Parse()
	log.Printf("server listening on port %d\n", port)

	tcpServer, err := server.NewTCPServer(port, maxCons, bufSize)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := tcpServer.Listener.Accept()

		if err != nil {
			log.Println(err)
			continue
		}

		tcpServer.Mut.Lock()
		tcpServer.Connections[conn.RemoteAddr().String()] = conn
		tcpServer.Mut.Unlock()

		log.Printf("client connected @ ip address %s", conn.RemoteAddr().String())
		go server.HandleClient(conn, tcpServer)
	}
}
