package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

type Packet struct {
	sourceAddr string
	destAddr   string
	data       string
}



func main() {
  logPath := "ghat.log"
  logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
  if err != nil {
    log.Fatal("can't open log file:", err)
    return
  }
  log.SetOutput(logFile)
	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()
	for {
		reader := bufio.NewReader((os.Stdin))
		fmt.Println("enter your chat: ")
		char, _, err := reader.ReadLine()
		if err != nil {
			log.Fatal(err)
			return
		}

		_, err = conn.Write([]byte(char))
		if err != nil {
			log.Fatal(err)
			return
		}
    
    buf := make([]byte, 2048)
    n, err := conn.Read(buf)
    if err != nil {
      log.Fatal(err)
      return
    }
    fmt.Println(string(buf[:n]))
	}
}
