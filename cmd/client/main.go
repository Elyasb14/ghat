package main

import (
    "bufio"
    "flag"
    "fmt"
    "net"
    "os"
    "github.com/Elyasb14/ghat/pkg/client"
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

    for {
        select {
        case msg := <- messages:
            fmt.Println(msg)
        }
    }
}
