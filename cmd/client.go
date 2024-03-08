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
	}
}
