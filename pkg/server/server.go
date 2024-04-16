package server

import (
  "fmt"
  "net"
)

type TCPServer struct {
  listener net.Listener
}

func NewTCPServer(port uint) (*TCPServer, error) {
  listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
  if err != nil {
    return nil, err
  } 
  
  return &TCPServer{
    listener: listener,
  }, nil
}

