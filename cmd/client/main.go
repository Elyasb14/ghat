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
	var addr string
	flag.StringVar(&addr, "addr", "localhost:8080", "addr to connect to")
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

	// for {
	//     select {
	//     case msg := <- messages:
	//         fmt.Println(msg)
	//         fmt.Print("> ")
	//     }
	// }

	// i am a little confused as to how this works
	// how does it know the channel is still open
	fmt.Println("remote messages are in" + "\033[31m" + " red " + "\033[0m" + "your messages are in white")
	for msg := range messages {
		fmt.Println(client.ColorString(msg))
		fmt.Print("> ")
	}
}
