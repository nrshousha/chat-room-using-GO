package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
	"strings"
)

type Message struct {
	Name string
	Text string
}

var messages []Message

func main() {
	client, err := rpc.Dial("tcp", "localhost:6700")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer client.Close()

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Printf("Welcome %s! You've joined the chat. Type a message to see the chat history.\n", name)

	for {
		fmt.Print("Enter message (or 'exit' to quit): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Exiting chat...")
			break
		}

		var reply []string
		msg := Message{Name: name, Text: input}

		err = client.Call("ChatServer.SendMessage", msg, &reply)
		if err != nil {
			fmt.Println("Error sending message:", err)
			break
		}

		fmt.Println("\n--- Chat History ---")
		for _, m := range reply {
			fmt.Println(m)
		}
		fmt.Println("--------------------")
	}
}
