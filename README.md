# Simple Chatroom in Go

This project is a small chatroom made with **Go**.
It has two files:

* **server.go** → runs the chat server
* **client.go** → lets users join the chat and send messages

You can open many terminals and run `client.go` in each one to chat between them.

---

## How it works

* The **server** listens on port `8080` and saves every message it receives.
* When a **client** sends a message, the server adds it to the chat and sends back the full chat history.
* The client then prints all messages so you can see what everyone said.

---

## Used Packages

All from Go’s standard library:

* `fmt` → prints text
* `log` → shows errors
* `net/http` → handles sending and receiving messages
* `encoding/json` → turns data into JSON
* `time` → adds timestamps to messages

---

## How to Run

1. Run the server:

   ```bash
   go run server.go
   ```
2. Open another terminal and run a client:

   ```bash
   go run client.go
   ```
3. Enter your name and start chatting!
4. Type `exit` to leave the chat.



https://github.com/user-attachments/assets/c7303959-ee35-4392-a4f2-70a8bf804ed1


