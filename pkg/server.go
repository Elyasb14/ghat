package server

import (
  "fmt"
  "net"
  "log"
)

type Message struct {
  Sender string
  Data string
}

func StartTCPServer() {
  ln, err := net.Listen("tcp", ":9999")
  log.Println("server init")
  if err != nil {
    fmt.Println(err)
    return
  }
  for {
    conn, err := ln.Accept()
    if err != nil {
      fmt.Println(err)
      continue
    }
    log.Println("connection accepted from", conn.RemoteAddr())
    
    go handleTCPConnection(conn)

  }
}

func handleTCPConnection(conn net.Conn) {
  defer conn.Close()
  
  for {
    buf := make([]byte, 2048)
    n, err := conn.Read(buf)
    if err != nil {
      fmt.Println(err)
      return
    }
    str := string(buf[:n])
    fmt.Println(str) 
  }
}
