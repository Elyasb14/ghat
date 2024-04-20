package server

import (
	"fmt"
	"net"
)

type Connection net.Conn

type TCPServer struct {
	Listener net.Listener
	Clients  []Connection
}

func NewTCPServer(port uint) (*TCPServer, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}

	return &TCPServer{
		Listener: listener,
		Clients:  make([]Connection, 0, 100),
	}, nil
}

func HandleClient(conn net.Conn, server *TCPServer) {}