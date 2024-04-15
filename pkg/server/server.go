package server

import (
  "fmt"
  "net"
)


func NewTCPServer(port uint16) (net.Listener, error) {
  listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
  if err != nil {
    return nil, err
  } 
  
  return listener, nil
}


