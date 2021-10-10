package main

import (
	"fmt"
	"net"
)

func server(serverListener net.Listener) {

}

func main() {
	serverListener, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println("Error initializing server: ", err)
		return
	}

	go server(serverListener)

	var (
		opt int = 0
	)

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
