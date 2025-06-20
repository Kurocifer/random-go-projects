package main

import (
	"bufio"
	"log"
	"net"
	"strings"
)

const (
	LISTENER_HOST = "0.0.0.0"
	LISTNER_PORT  = "8080"
	CONN_TYPE     = "tcp"
)

func main() {
	listernerAddr := LISTENER_HOST + ":" + LISTNER_PORT
	log.Println("Startcin TCP echo server on " + listernerAddr)

	listener, err := net.Listen(CONN_TYPE, listernerAddr)
	if err != nil {
		log.Printf("Error: can't listen - %s\n", err.Error())
	}
	defer listener.Close()

	log.Println("Listening for incoming connectins...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error: Accepting connectiong - %s\n", err.Error())
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	clientAddr := conn.RemoteAddr().String()
	log.Printf("New client connected %s\n", clientAddr)

	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			// Client disconnected
			if err.Error() == "EOF" {
				log.Printf("Client %s disconnected\n", clientAddr)
				return
			}

			log.Printf("An error occured while reading client %s message: %s\n", clientAddr, err.Error())
			return
		}

		trimmedMessage := strings.TrimSpace(message)
		log.Printf("Message recieved from %s: %s\n", clientAddr, trimmedMessage)

		_, err = conn.Write([]byte(message))
		if err != nil {
			log.Printf("An error occured while echoing message to client: %s", err.Error())
			return
		}

		log.Println("Message echoed to client")
	}
}
