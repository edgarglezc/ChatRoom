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

func (r *Request) Show() string {
	var msg string
	switch r.Type {
	case CONNECTION:
		msg = r.Client + " has arrived to the ChatRoom!"
	case DISCONNECTION:
		msg = r.Client + " has disconnected from the ChatRoom!"
	case MESSAGE:
		msg = r.Client + ": " + r.Message
	case FILE:
		msg = r.Client + " has sent a file: " + r.Message
	case END:
		msg = r.Message
	}
	return msg
}

const (
	CONNECTION    int = 1
	DISCONNECTION     = 2
	MESSAGE           = 3
	FILE              = 4
	END               = 5
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
		fmt.Println("Welcome to the ChatRoom, " + name + "!")
		fmt.Println("[1] Send message")
		fmt.Println("[2] Send file")
		fmt.Println("[3] Show messages")
		fmt.Println("[4] Exit chat room")
		fmt.Scanln(&opt)

		switch opt {
		case 1:
			sendMessage(clientDial, name, &requests)
		case 2:
			sendFile(clientDial, name, &requests)
		case 3:
			showMessages(&requests)
		case 4:
			disconnection(clientDial, name)
			fmt.Println("See you soon!")
		default:
			fmt.Println("Option not found")
		}
	}

	clientDial.Close()
}

func client(clientDial net.Conn, requests *[]Request, name string) {
	for {
		var request Request

		err := gob.NewDecoder(clientDial).Decode(&request)
		if err != nil {
			fmt.Println("Error decoding request: ", err)
			continue
		}

		if request.Type == END {
			fmt.Println(request.Show())
			clientDial.Close()
			os.Exit(0)
		}

		*requests = append(*requests, request)
		if request.Client != name {
			fmt.Println(request.Show())
		}
	}
}

func connection(clientDial net.Conn, name string) {
	r := Request{
		Type:   CONNECTION,
		Client: name,
	}
	err := gob.NewEncoder(clientDial).Encode(r)
	if err != nil {
		fmt.Println("Error sending request: ", err)
	}
}

func sendMessage(clientDial net.Conn, name string, requests *[]Request) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Write something: ")
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

func sendFile(clientDial net.Conn, name string, requests *[]Request) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter filename: ")
	scanner.Scan()
	fileName := scanner.Text()

	fileData, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("An error has ocurred trying to open the file: ", err)
		return
	}

	r := Request{
		Type:    FILE,
		Client:  name,
		Message: fileName,
		Data:    fileData,
	}

	err = gob.NewEncoder(clientDial).Encode(r)
	if err != nil {
		fmt.Println("Error sending file: ", err)
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

func showMessages(requests *[]Request) {
	fmt.Println("------------------")
	for _, v := range *requests {
		fmt.Println(v.Show())
	}
	fmt.Println("------------------")
}
