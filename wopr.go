package main

import (
	"fmt"
	"net"
	"strings"
)

var clients map[net.Conn]clientData

type clientData struct{ loggedIn bool }

func main() {
	listener, err := net.Listen("tcp", ":23")
	if err != nil {
		panic(err)
	}

	fmt.Println("War Operation Plan Response (WOPR) has initialized.")

	for {
		client, err := listener.Accept()
		if err != nil {
			continue
		}

		go connectClient(client)
	}
}

func connectClient(client net.Conn) {
	client.Write([]byte("LOGON: "))

	for {
		buf := make([]byte, 4096)
		numbytes, err := client.Read(buf)
		if numbytes == 0 || err != nil {
			return
		}
		handleInput(string(buf), client)
	}
}

func handleInput(input string, client net.Conn) {

	fmt.Println("handleInput:", input)

	if strings.Contains(input, "Joshua") {
		clients[client].loggedIn = true
		sendString(client, "GREETING PROFESSOR FALKEN.\n")
	}
}

func sendString(client net.Conn, msg string) {
	client.Write([]byte(msg))
}
