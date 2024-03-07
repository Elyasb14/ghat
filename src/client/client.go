package main 

import (
  "net"
  "fmt"
)


func main() {
  conn, err := net.Dial("tcp", "localhost:9999")
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
