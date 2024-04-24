package server

import (
	"log"
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
  for {
	  n, err := conn.Read(buf)
    if err != nil {
			conn.Close()
			log.Printf("client @ ip addr %s disconnected", conn.RemoteAddr().String())
			return
    }
   
    text := string(buf[0:n])
    conn.Write([]byte(text))
  }
}
