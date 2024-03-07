package server

import (
  "fmt"
  "net"
  "log"
)

func StartTCPServer() {
  ln, err := net.Listen("tcp", ":9999")
  if err != nil {
    fmt.Println(err)
    return
  }
  for {
    conn, err := ln.Accept()
    fmt.Println("connection accepted")
    if err != nil {
      fmt.Println(err)
      continue
    }

    go handleTCPConnection(conn)

  }
}

func handleTCPConnection(conn net.Conn) {
  defer conn.Close()
  
  for {
    buf := make([]byte, 1024)
    n, err := conn.Read(buf)
    if err != nil {
      log.Fatal(err)
      return 
    }
    str := string(buf[:n])
    fmt.Println(str) 
  }
}
