package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Message struct {
	Author  string    `json:"author"`
	Content string    `json:"content"`
	Time    time.Time `json:"time"`
}

var messages []Message

func sendHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var msg Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, "Invalid message format", http.StatusBadRequest)
		return
	}

	msg.Time = time.Now()
	messages = append(messages, msg)
	fmt.Printf("[%s] %s: %s\n", msg.Time.Format("15:04:05"), msg.Author, msg.Content)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func historyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func main() {
	http.HandleFunc("/send", sendHandler)
	http.HandleFunc("/history", historyHandler)

	fmt.Println("🚀 Chat server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
