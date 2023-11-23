package main

import (
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

	conn.Write([]byte("hello world!"))

	incoming_msg := make([]byte, 64)
	conn.Read(incoming_msg)

	fmt.Println("incoming message:", incoming_msg)
}
