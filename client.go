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

func client(clientDial net.Conn, clients map[string]net.Conn, requests *[]Request) {

}

func test(messages *[]*string) {
	for i := 0; i < 3; i++ {
		var str string
		str = "Hola"
		*messages = append(*messages, &str)
	}
}

func main() {
	clientDial, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println("Error initializing client: ", err)
		return
	}

	opt := 0
	clients := make(map[string]net.Conn)
	requests := make([]Request, 0)

	go client(clientDial, clients, &requests)

	for opt != 4 {
		fmt.Println("Welcome to the ChatRoom!")
		fmt.Println("[1] Send message")
		fmt.Println("[2] Send file")
		fmt.Println("[3] Show message")
		fmt.Println("[4] Exit chat room")
		fmt.Print("=> ")
		fmt.Scanln(&opt)

		switch opt {
		case 1:
			fmt.Println("enviando mensaje")
		case 2:
			fmt.Println("enviando archivo")
		case 3:
			fmt.Println("mostrando mensajes")
		case 4:
			fmt.Println("Discconected")
		default:
			fmt.Println("Option not found")
		}
	}
}
