package main

import (
	"fmt"
	"net"
)

type File struct {
	extension string
	data      []byte
}

func client(clientDia net.Conn, messages *[]*string, files *[]*File) {
	for i := 0; i < 3; i++ {
		var str *string
		*str = "Hola"
		*messages = append(*messages, str)
	}
}

func main() {
	clientDial, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println("Error initializing client: ", err)
	}

	var (
		opt      int = 0
		messages *[]*string
		files    *[]*File
	)

	go client(clientDial, messages, files)

	for _, v := range *messages {
		fmt.Println(v)
	}

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
