package server

import (
	"log"
  "fmt"
	"net"
	"os"
)

// this gets called implicitly when the server package gets called
func init() {
	logPath := "ghat.log"

	logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Println("can't open log file:", err)
		return
	}
	log.SetOutput(logFile)
}

func StartTCPServer() {

	ln, err := net.Listen("tcp", ":9999")
	log.Println("server init")
	if err != nil {
		log.Println(err)
		return
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
		 log.Println(err)
			continue
		}
		log.Println("connection accepted from", conn.RemoteAddr())

		go handleTCPConnection(conn)

	}
}

func handleTCPConnection(conn net.Conn) {
	defer conn.Close()
  fmt.Println("you are connected to the server")
	for {
		buf := make([]byte, 2048)
		n, err := conn.Read(buf)
		if err != nil {
			log.Println(err)
			return
		}
		//str := string(buf[:n])
		//fmt.Println(str)
    conn.Write(buf[:n])
	}
}
