package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

type Request struct {
	Type    int
	Client  string
	Message string
	Data    []byte
}

const (
	CONNECTION    int = 1
	DISCONNECTION     = 2
	MESSAGE           = 3
	FILE              = 4
)

func main() {
	serverListener, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println("Error initializing server: ", err)
		return
	}

	clients := make(map[string]net.Conn)
	opt := 0

	go server(serverListener, clients)

	for opt != 3 {
		fmt.Println("ChatRoom Server Dashboard")
		fmt.Println("[1] Show messages/files")
		fmt.Println("[2] Backup messages/files")
		fmt.Println("[3] End server")
		fmt.Print("=> ")
		fmt.Scanln(&opt)

		switch opt {
		case 1:
			fmt.Println("mensajes")
		case 2:
			fmt.Println("respaldo mensajes")
		case 3:
			fmt.Println("se terminó el server")
		default:
			fmt.Println("Option not found")
		}
	}
}

func server(serverListener net.Listener, clients map[string]net.Conn) {
	for {
		client, err := serverListener.Accept()
		if err != nil {
			fmt.Println("Error connecting with client: ", err)
			continue
		}
		go handleClient(client, clients)
	}
}

func handleClient(client net.Conn, clients map[string]net.Conn) {
	for {
		var request Request
		err := gob.NewDecoder(client).Decode(&request)
		if err != nil {
			fmt.Println("Error decoding request: ", err.Error())
			continue
		}
		switch request.Type {
		case CONNECTION:
			clients[request.Client] = client
			fmt.Printf("%s has arrived to the ChatRoom!\n", request.Client)
		case DISCONNECTION:
			delete(clients, request.Client)
			fmt.Printf("%s has disconnected from the ChatRoom!\n", request.Client)
		case MESSAGE:
		case FILE:
		default:
			fmt.Println("An error has ocurred...")
		}
	}
}
