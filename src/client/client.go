package client

import (
  "net"
  "fmt"
)


func ClientHandler() {
  conn, err := net.Dial("tcp", "localhost:9999")
  defer conn.Close()
  if err != nil {
    fmt.Println(err)
    return
  }
  _, err = conn.Write([]byte("hello server, from the client"))
  if err != nil {
    fmt.Println(err)
    return
  }
}
