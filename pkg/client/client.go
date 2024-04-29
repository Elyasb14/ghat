package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func ReadFromServer(conn net.Conn, messages chan string) {
	buf := make([]byte, 64)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("can't read from server, crashing client")
			os.Exit(0)
		}

		text := string(buf[0:n])
		messages <- text
	}
}

func ReadFromUser(conn net.Conn, reader *bufio.Reader, messages chan string) {
	// read from user input
	for {
		fmt.Print("> ")
		// newline is delimeter to read until
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		// indexing is so you don't include a newline char
		conn.Write([]byte(input[:len(input)-1]))
	}
}
