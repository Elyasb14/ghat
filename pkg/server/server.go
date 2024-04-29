package server

import (
	"fmt"
	"log"
	"net"
	"sync"
)

type Connection net.Conn

type TCPServer struct {
	Listener    net.Listener
	Connections map[string]Connection
	Mut         sync.Mutex
}

func colorText(text string) string {
	return "\033[31m" + text + "\033[0m"
}

func NewTCPServer(port uint) (*TCPServer, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}

	return &TCPServer{
		Listener:    listener,
		Connections: make(map[string]Connection, 0),
	}, nil
}

func HandleClient(conn net.Conn, server *TCPServer) {
	buf := make([]byte, 64)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			conn.Close()
			log.Printf("client @ ip addr %s disconnected", conn.RemoteAddr().String())
			server.Mut.Lock()
			delete(server.Connections, conn.RemoteAddr().String())
			server.Mut.Unlock()
			return
		}

		text := colorText(string((buf[0:n])))
		log.Printf("message: %s from ip addr %s", text, conn.RemoteAddr().String())

		server.Mut.Lock()
		for _, client := range server.Connections {
			if client == conn {
				continue
			}

			if text == "\n" || text == " " {
				continue
			}

			_, err := client.Write([]byte(text))
			server.Mut.Unlock()
			if err != nil {
				log.Println(err)
				continue
			}
		}

	}
}
