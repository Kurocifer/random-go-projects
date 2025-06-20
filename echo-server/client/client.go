package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const SERVER_ADDR = "localhost:8080"

func main() {

	conn, err := net.Dial("tcp", SERVER_ADDR)
	if err != nil {
		log.Printf("Error: Could not establish connectionn with server - %s\n", err.Error())
	}
	defer conn.Close()

	log.Printf("Established connection with server at %s\n", SERVER_ADDR)
	log.Printf("Enter text to send to server\n")

	go func() {
		reader := bufio.NewReader(conn)

		for {
			message, err := reader.ReadString('\n')
			if err != nil {
				if err.Error() == "EOF" {
					log.Printf("Server disconnected\n")
					return
				}
				log.Printf("Error while reading from server: %s\n", err.Error())
			}

			message = strings.TrimSpace(message)
			log.Printf("Server echo: %s\n", message)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()

		_, err := fmt.Fprintf(conn, "%s\n", text)
		if err != nil {
			log.Printf("Error while sending message to server: %s", err.Error())
		}
	}
	if err := scanner.Err(); err != nil {
		log.Printf("Error could not read input: %s", err.Error())
	}

	log.Printf("Exiting...")
}
