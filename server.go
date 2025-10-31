package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
)

type ChatServer struct {
	messages []string
}

type Message struct {
	Name string
	Text string
}

func (c *ChatServer) SendMessage(msg Message, reply *[]string) error {
	if msg.Text == "" {
		return errors.New("empty message")
	}

	formatted := fmt.Sprintf("%s: %s", msg.Name, msg.Text)
	c.messages = append(c.messages, formatted)

	fmt.Println(formatted)

	*reply = c.messages
	return nil
}

func main() {
	chatServer := new(ChatServer)
	rpc.Register(chatServer)

	listener, err := net.Listen("tcp", ":6700")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Chat server running on port 6700...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
