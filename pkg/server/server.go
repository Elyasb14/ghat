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
	bufSize     uint
	maxCons     uint
	// Messages    chan string
}

func NewTCPServer(port uint, maxCons uint, bufSize uint) (*TCPServer, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}

	return &TCPServer{
		Listener:    listener,
		Connections: make(map[string]Connection, maxCons),
		bufSize:     bufSize,
		maxCons:     maxCons,
	}, nil
}

func HandleClient(conn net.Conn, server *TCPServer) {
	buf := make([]byte, server.bufSize)

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

		text := string((buf[0:n]))
		log.Printf("message: %s from ip addr %s", text, conn.RemoteAddr().String())

		go BroadCastMessage(server, conn, text)
	}
}

func BroadCastMessage(server *TCPServer, conn net.Conn, text string) {
	server.Mut.Lock()
	for _, client := range server.Connections {
		if client == conn {
			continue
		}

		if text == "\n" || text == " " {
			continue
		}

		_, err := client.Write([]byte(text))
		if err != nil {
			log.Println(err)
			continue
		}
	}
	server.Mut.Unlock()
}
