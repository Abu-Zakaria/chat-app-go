package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"slices"
)

func main() {
	registered_users := []string{}

	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic("Couldn't listen to the port!")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Something went wrong while accepting an incoming connection")
			continue
		}

		defer conn.Close()

		remote_addr := conn.RemoteAddr()

		if !slices.Contains(registered_users, remote_addr.String()) {
			registered_users = append(registered_users, remote_addr.String())
		}

		log.Println("Received a connection from:", remote_addr)

		msg := make([]byte, 64)

		length, err := conn.Read(msg)
		if err != nil {
			log.Println("Failed to read the incoming message")
		}
		log.Println("incoming message length: ", length)
		log.Println("incoming message: ", string(msg))

		fmt.Println("registered:", registered_users)

		for _, address := range registered_users {
			sendMessage(conn, address, msg)
		}
	}
}

func sendMessage(conn net.Conn, address string, message []byte) error {
	length, err := conn.Write(message)
	if err != nil {
		return errors.New("sendMessage(): Something went sending message to:" + address)
	}
	if length == 0 {
		return errors.New("sendMessage(): The message was invalid")
	}

	return nil
}
