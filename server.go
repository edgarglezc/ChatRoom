package main

import (
	"fmt"
	"net"
)

type File struct {
	name      string
	extension string
	data      []byte
}

func server(serverListener net.Listener, messages *[]*string, files *[]*File, clients *[]*net.Conn) {
	for {
		client, err := serverListener.Accept()
		if err != nil {
			fmt.Println("Error connecting with client: ", err)
			continue
		}
		go handleClient(client, messages, files, clients)
	}
}

func handleClient(client net.Conn, messages *[]*string, files *[]*File, clients *[]*net.Conn) {

}

func showMessagesAndFiles(messages *[]*string, files *[]*File) {

}

func backupMessagesAndFiles(messages *[]*string, files *[]*File) {

}

func main() {
	serverListener, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println("Error initializing server: ", err)
		return
	}

	var (
		opt      int = 0
		messages []*string
		files    []*File
		clients  []*net.Conn
	)

	go server(serverListener, &messages, &files, &clients)

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
