package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"net-cat/functions"
)

var (
	port             string = "8989"
	errProgram       error  = errors.New("[USAGE]: ./TCPChat $port")
	clients                 = make(map[string]net.Conn)
	connectedClients        = make(chan struct{}, 2)
	messageHistory   string
)

func main() {
	// check arguments and get the port
	if len(os.Args) > 2 {
		fmt.Println(errProgram)
		return
	} else if len(os.Args) == 2 {
		if os.Args[1]== "" {
			fmt.Println(errProgram)
			return
		}
		port = os.Args[1]

	}
	

	// Create a TCP listener (server)
	server, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer server.Close()

	fmt.Printf("Listening on the port :%s\n", port)

	// Accept incoming connection
	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle the connection in seperate goroutine
		go functions.HandleConnection(conn, clients, connectedClients, &messageHistory)
	}
}
