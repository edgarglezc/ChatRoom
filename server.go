package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"time"
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
	serverListener, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println("Error initializing server: ", err)
		return
	}

	clients := make(map[string]net.Conn)
	requests := make([]Request, 0)
	opt := 0

	go server(serverListener, clients, &requests)

	for opt != 3 {
		fmt.Println("ChatRoom Server Dashboard")
		fmt.Println("[1] Show messages/files")
		fmt.Println("[2] Backup messages/files")
		fmt.Println("[3] End server")
		fmt.Scanln(&opt)

		switch opt {
		case 1:
			showRequests(&requests)
		case 2:
			backupRequests(clients, &requests)
			fmt.Println("Backup created!")
		case 3:
			endServer(clients)
		default:
			fmt.Println("Option not found")
		}
	}
}

func server(serverListener net.Listener, clients map[string]net.Conn, requests *[]Request) {
	for {
		client, err := serverListener.Accept()
		if err != nil {
			fmt.Println("Error connecting with client: ", err)
			continue
		}
		go handleClient(client, clients, requests)
	}
}

func handleClient(client net.Conn, clients map[string]net.Conn, requests *[]Request) {
	for {
		var request Request
		err := gob.NewDecoder(client).Decode(&request)
		if err != nil {
			fmt.Println("Error decoding request: ", err.Error())
			continue
		}

		if request.Type == CONNECTION {
			clients[request.Client] = client
		}

		if request.Type == DISCONNECTION {
			delete(clients, request.Client)
			*requests = append(*requests, request)
			sendRequest(client, clients, request)
			fmt.Println(request.Show())
			return
		}

		*requests = append(*requests, request)
		sendRequest(client, clients, request)
		fmt.Println(request.Show())
	}
}

func showRequests(requests *[]Request) {
	fmt.Println("------------------")
	for _, v := range *requests {
		fmt.Println(v.Show())
	}
	fmt.Println("------------------")
}

func backupRequests(clients map[string]net.Conn, requests *[]Request) {
	_, err := os.Stat("backup")
	if os.IsNotExist(err) {
		os.Mkdir("backup", 0755)
	}

	file, err := os.Create("./backup/backup.txt")
	if err != nil {
		fmt.Println("Error creating backup file: ", err)
		return
	}
	defer file.Close()

	for _, v := range *requests {
		file.WriteString(v.Show() + "\n")
	}

	backupFiles(clients, requests)
}

func backupFiles(clients map[string]net.Conn, requests *[]Request) {
	_, err := os.Stat("received_files")
	if os.IsNotExist(err) {
		os.Mkdir("received_files", 0755)
	}

	for id, _ := range clients {
		clientDirPath := "./received_files/" + id
		_, err := os.Stat(clientDirPath)
		if os.IsNotExist(err) {
			os.Mkdir(clientDirPath, 0755)
		}
	}

	for id, _ := range clients {
		clientDirPath := "./received_files/" + id
		for _, v := range *requests {
			if v.Client != id && v.Type == FILE {
				date := getDate()
				file, err := os.Create(clientDirPath + "/" + v.Client + "_" + date + "_" + v.Message)
				if err != nil {
					fmt.Println("Error creating request file: ", err)
					continue
				}
				defer file.Close()
				file.WriteString(string(v.Data))
			}
		}
	}
}

func getDate() string {
	time := time.Now().Format("01022006150405")
	return time
}

func sendRequest(client net.Conn, clients map[string]net.Conn, request Request) {
	for _, conn := range clients {
		err := gob.NewEncoder(conn).Encode(&request)
		if err != nil {
			fmt.Println("Error encoding request: ", err)
			continue
		}
	}
}

func endServer(clients map[string]net.Conn) {
	r := Request{
		Type:    END,
		Message: "The session was ended forcefully by the server",
	}

	for id, conn := range clients {
		err := gob.NewEncoder(conn).Encode(&r)
		if err != nil {
			fmt.Println("Error ending session with client: ", err)
			continue
		}
		delete(clients, id)
		defer conn.Close()
	}

	fmt.Println("The ChatRoom has ended!")
	os.Exit(0)
}
