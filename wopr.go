package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

var clients = make(map[net.Conn]bool)

func main() {

	port := "2000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}

	fmt.Println("War Operation Plan Response (WOPR) has initialized.\nRunning on port: " + port)

	for {
		client, err := listener.Accept()
		if err != nil {
			continue
		}

		go connectClient(client)
	}
}

func connectClient(client net.Conn) {

	// Clear Screen
	sendString(client, "\u001B[2J")
	sendString(client, "LOGON: ")

	for {
		buf := make([]byte, 4096)
		numBytes, err := client.Read(buf)
		if numBytes == 0 || err != nil {
			return
		}

		// If we get telnet commands, ignore them
		if buf[0] == 255 {
			continue
		}

		handleInput(string(buf[0:numBytes]), client)
	}
}

func handleInput(input string, client net.Conn) {

	input = stripInput(input)
	fmt.Println("handleInput:", input)

	if !clientLoggedIn(client) {
		disconnect := true

		switch input {
		case "JOSHUA":
			loginClient(client)
			sendString(client, "GREETING PROFESSOR FALKEN.\r\n")
			disconnect = false
		case "HELP", "HELP LOGON":
			sendString(client, "HELP NOT AVAILABLE.\r\n")
		default:
			sendString(client, "IDENTIFICATION NOT RECOGNIZED BY SYSTEM.\r\n")
		}

		if disconnect {
			sendString(client, "--CONNECTION TERMINATED--\r\n")
			client.Close()
		}
	} else {
		switch input {
		case "EXIT":
			client.Close()
		case "HELLO":
			sendString(client, "HOW ARE YOU FEELING TODAY?\r\n")
		case "I AM FINE. HOW ARE YOU?":
			sendString(client, "EXCELLENT. IT'S BEEN A LONG TIME.\r\n")
			sendString(client, "CAN YOU EXPLAIN THE REMOVAL OF YOUR USER ACCOUNT ON 6/23/73?\r\n")
		case "HELP GAMES":
			sendString(client, "'GAMES' REFERS TO MODELS, SIMULATIONS AND GAMES WHICH HAVE")
			sendString(client, " TACTICAL AND STRATEGIC APPLICATIONS\r\n")
		case "LIST GAMES":
			listGames(client)
		case "HELP SYSTEM":
			sendString(client, "HELP GAMES\r\nLIST GAMES\r\nEXIT\r\n")
		case "PEOPLE SOMETIMES MAKE MISTAKES":
			sendString(client, "YES THEY DO. SHALL WE PLAY A GAME?\r\n")
		case "LOVE TO. HOW ABOUT GLOBAL THERMONUCLEAR WAR":
			sendString(client, "WOULDN'T YOU PREFER A GOOD GAME OF CHESS?\r\n")
		case "LATER. LET'S PLAY GLOBAL THERMONUCLEAR WAR":
			sendString(client, "FINE\r\n")
		case "":
			// Do nothing
		default:
			sendString(client, "UNKNOWN COMMAND: '"+input+"'\r\n")
		}
	}
}

func listGames(client net.Conn) {
	sendString(client, "FALKEN'S MAZE\r\n")
	sendString(client, "BLACK JACK\r\n")
	sendString(client, "GIN RUMMY\r\n")
	sendString(client, "HEARTS\r\n")
	sendString(client, "BRIDGE\r\n")
	sendString(client, "CHECKERS\r\n")
	sendString(client, "CHESS\r\n")
	sendString(client, "POKER\r\n")
	sendString(client, "FIGHTER COMBAT\r\n")
	sendString(client, "GUERILLA ENGAGEMENT\r\n")
	sendString(client, "DESERT WARFARE\r\n")
	sendString(client, "AIR-TO-GROUND ACTIONS\r\n")
	sendString(client, "THEATERWIDE TACTICAL WARFARE\r\n")
	sendString(client, "THEATERWIDE BIOTOXIC AND CHEMICAL WARFARE\r\n\r\n")
	sendString(client, "GLOBAL THERMONUCLEAR WAR\r\n")
}

func loginClient(client net.Conn) {
	clients[client] = true
}

func clientLoggedIn(client net.Conn) bool {
	return clients[client]
}

func sendString(client net.Conn, msg string) {
	client.Write([]byte(msg))
}

func stripInput(input string) string {
	input = strings.Replace(input, "\r", "", -1)
	input = strings.Replace(input, "\n", "", -1)
	input = strings.ToUpper(input)

	return input
}
