package main

import (
    "bufio"
    "flag"
    "fmt"
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

    // Create a new reader for reading from the standard input
    reader := bufio.NewReader(os.Stdin)

    // Create a new channel for receiving messages from the server
    messages := make(chan string)

    // Start a new goroutine to read messages from the server
    go func() {
        buf := make([]byte, 1024)
        for {
            n, err := conn.Read(buf)
            if err != nil {
                fmt.Println("Error reading from server:", err)
                return
            }
            messages <- string(buf[:n])
        }
    }()

    // Start a new goroutine to read user input
    go func() {
        for {
            fmt.Print("> ")
            input, err := reader.ReadString('\n')
            if err != nil {
                fmt.Println("Error reading input:", err)
                continue
            }

            // Remove the newline character from the input
            input = input[:len(input)-1]

            // Send the input to the server
            _, err = conn.Write([]byte(input))
            if err != nil {
                fmt.Println("Error sending message:", err)
                continue
            }
        }
    }()

    // Continuously check for messages from the server
    for {
        select {
        case msg := <-messages:
            fmt.Println(msg)
        }
    }
}