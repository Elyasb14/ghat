package server

import (
	"fmt"
	"net"
)

type TCPServer struct {
	Listener net.Listener
}

func NewTCPServer(port uint) (*TCPServer, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}

	return &TCPServer{
		Listener: listener,
	}, nil
}

func HandleClient(conn net.Conn) {
	buf := make([]byte, 64)
	conn.Read(buf)

}