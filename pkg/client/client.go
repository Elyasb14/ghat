package client

import (
	"bufio"
	"fmt"
	"net"
)

func ReadFromServer(conn net.Conn, messages chan string) {
	buf := make([]byte, 64)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
		}

		text := string(buf[0:n])
		messages <- text
	}
}

func ReadFromUser(conn net.Conn, reader *bufio.Reader, messages chan string) {
	// read from user input
	for {
		input, err := reader.ReadString('\n') // newline is delimeter to read until
		if err != nil {
			fmt.Println(err)
		}
		conn.Write([]byte(input))
	}

}
