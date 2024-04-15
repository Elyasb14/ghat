package main

import (
  "github.com/Elyasb14/ghat/pkg/server"
  "fmt"
)



func main() {
  server, err := server.NewTCPServer(8080)
  if err != nil {
    fmt.Println(err)
  } 
  
  for {
    server.listener 
  }
}
