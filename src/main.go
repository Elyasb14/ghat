package main

import (
  "ebianchi/ghat/src/server"
)

type Message struct {
	Sender string
	Data   string
}

func main() {
	server.StartTCPServer()
}
