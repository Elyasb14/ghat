package server

import (
	"log"
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

type Packet struct {
  sender []byte
  message []byte
}

func StartTCPServer() {

	ln, err := net.Listen("tcp", ":9999")
	log.Println("server init")
	if err != nil {
		log.Println(err)
		return
	}
  
  // make a channel I can send packets through, make a hashmap of clients
  // pass in the channel and clients to the broadcast go routine
  messageChan := make(chan Packet)
  clients := make(map[net.Conn]struct{})
  go broadcastMessages(messageChan, clients)

	for {
		conn, err := ln.Accept()
		if err != nil {
		 log.Println(err)
			continue
		}
		log.Println("connection accepted from:", conn.RemoteAddr())
     
    // add the new connection to the client hashmap, this will get passed into the 
    // broadcastMessage goroutine
    clients[conn] = struct{}{}

		go handleTCPConnection(conn, messageChan)

	}
}

func handleTCPConnection(conn net.Conn, messageChan chan<- Packet) {
	defer func() {
    conn.Close()
    // delete(clients, conn)
  }()
	for {
		buf := make([]byte, 2048)
		n, err := conn.Read(buf)
    packet := Packet{
      sender:  []byte("from " + conn.RemoteAddr().String() + ": "),
      message: buf[:n],
    }
		if err != nil {
			log.Println(err)
			return
		}

    log.Printf("message: %s from client @ ip addr: %s", buf[:n], conn.RemoteAddr())
    //conn.Write(packet.sender) 
    //conn.Write(packet.message)
    //send the packet to the message channel
    messageChan <- packet
	}
}

func broadcastMessages(messageChan <- chan Packet, clients map[net.Conn]struct{}) {
  for {
    // recieve packet from messageChan, save it into packet, nice syntax!
    packet := <- messageChan
    for conn := range clients {
      conn.Write(append(packet.sender, packet.message...))
    }
  }
}

