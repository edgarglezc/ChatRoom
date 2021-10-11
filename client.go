package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
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
	clientDial, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println("Error initializing client: ", err)
		return
	}

	var name string
	opt := 0
	requests := make([]Request, 0)

	fmt.Print("Username: ")
	fmt.Scanln(&name)

	connection(clientDial, name)
	go client(clientDial, &requests, name)

	for opt != 4 {
		fmt.Println("Welcome to the ChatRoom!")
		fmt.Println("[1] Send message")
		fmt.Println("[2] Send file")
		fmt.Println("[3] Show messages")
		fmt.Println("[4] Exit chat room")
		fmt.Print("=> ")
		fmt.Scanln(&opt)

		switch opt {
		case 1:
			sendMessage(clientDial, name)
		case 2:
			fmt.Println("enviando archivo")
		case 3:
			fmt.Println("mostrando mensajes")
		case 4:
			disconnection(clientDial, name)
		default:
			fmt.Println("Option not found")
		}
	}

	clientDial.Close()
}

func client(clientDial net.Conn, requests *[]Request, name string) {
	for {

	}
}

func connection(clientDial net.Conn, name string) {
	r := Request{
		Type:   1,
		Client: name,
	}
	err := gob.NewEncoder(clientDial).Encode(r)
	if err != nil {
		fmt.Println("Error sending request: ", err)
	}
}

func sendMessage(clientDial net.Conn, name string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Write something:")
	scanner.Scan()
	msg := scanner.Text()

	r := Request{
		Type:    MESSAGE,
		Client:  name,
		Message: msg,
	}

	err := gob.NewEncoder(clientDial).Encode(r)
	if err != nil {
		fmt.Println("Error sending message: ", err)
	}
}

func disconnection(clientDial net.Conn, name string) {
	request := Request{
		Type:   DISCONNECTION,
		Client: name,
	}
	err := gob.NewEncoder(clientDial).Encode(request)
	if err != nil {
		fmt.Println("Error sending request: ", err)
		return
	}
}
