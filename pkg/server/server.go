package server

import (
	"fmt"
	"net"
)

type Connection net.Conn

type TCPServer struct {
	listener net.Listener
	clients  []Connection
}

func NewTCPServer(port uint) (*TCPServer, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}

	return &TCPServer{
		listener: listener,
		clients:  make([]Connection, 0, 100),
	}, nil
}
