package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Message struct {
	Author  string    `json:"author"`
	Content string    `json:"content"`
	Time    time.Time `json:"time"`
}

func main() {
	fmt.Print("Enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Println("Type messages (type 'exit' to quit):")

	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "" {
			continue
		}

		if strings.ToLower(text) == "exit" {
			fmt.Println("👋 Exiting chat.")
			break
		}

		msg := Message{
			Author:  name,
			Content: text,
			Time:    time.Now(),
		}

		body, _ := json.Marshal(msg)
		resp, err := http.Post("http://localhost:8080/send", "application/json", bytes.NewBuffer(body))
		if err != nil {
			log.Println("❌ Failed to send message:", err)
			time.Sleep(2 * time.Second)
			continue
		}
		defer resp.Body.Close()

		var history []Message
		json.NewDecoder(resp.Body).Decode(&history)

		fmt.Println("\n--- Chat History ---")
		for _, m := range history {
			fmt.Printf("[%s] %s: %s\n", m.Time.Format("15:04"), m.Author, m.Content)
		}
		fmt.Println("--------------------")
	}
}
