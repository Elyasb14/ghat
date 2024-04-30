package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/Elyasb14/ghat/pkg/client"
	"net"
	"os"
)

func main() {
	var addr, color string
	flag.StringVar(&addr, "addr", "localhost:8080", "addr to connect to")
	flag.StringVar(&color, "color", "red", "color to print remote clients messages as")
	flag.Parse()

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	messages := make(chan string)

	go client.ReadFromUser(conn, reader, messages)
	go client.ReadFromServer(conn, messages)

	fmt.Println(client.ColorString("messages from remote clients will appear in the color this is printed in", color))
	for msg := range messages {
		fmt.Println(client.ColorString(msg, color))
		fmt.Print("> ")
	}
}
