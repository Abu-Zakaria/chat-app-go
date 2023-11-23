package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Please add the server address!")
		return
	}

	server_address := args[1]

	conn, err := net.Dial("tcp", server_address)
	if err != nil {
		log.Println("Couldn't connect with the server")
		log.Println("ERROR:", err)
	}

	defer conn.Close()

	for {
		fmt.Print("> ")

		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
			continue
		}

		conn.Write([]byte(text))
		log.Println("sent message!")

		incoming_msg := make([]byte, 64)
		conn.Read(incoming_msg)

		fmt.Println("incoming message:", incoming_msg)
	}
}
