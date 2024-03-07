package main

import (
  "ebianchi/ghat/src/server"
  "ebianchi/ghat/src/client"
  "log"
  "os"
)

type Message struct {
	Sender string
	Data   string
}

func main() {
  logFile, err := os.OpenFile("./ghat.log", os.O_WRONLY, 0644)
  if err != nil {
    log.Fatal("couldn't open log file")
  }
  defer logFile.Close()
	server.StartTCPServer()
  go client.ClientHandler()
}
