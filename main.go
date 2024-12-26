// Copyright © 2024 Mohammad Shafay Joyo
package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var connections = make(map[*websocket.Conn]bool)
var messageHistory []string
var mu sync.Mutex // To protect messageHistory

func main() {
	http.HandleFunc("/ws", handleConnections)
	http.Handle("/", http.FileServer(http.Dir("./")))

	fmt.Println("Server started at :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}
	defer conn.Close()

	connections[conn] = true

	// Send message history to the newly connected user
	mu.Lock()
	for _, msg := range messageHistory {
		conn.WriteMessage(websocket.TextMessage, []byte(msg))
	}
	mu.Unlock()

	for {
		var msg string
		err := conn.ReadMessage(&msg)
		if err != nil {
			fmt.Println("Error reading message:", err)
			delete(connections, conn)
			break
		}

		// Append message to history
		mu.Lock()
		messageHistory = append(messageHistory, msg)
		mu.Unlock()

		// Broadcast the message to all connected clients
		broadcastMessage(msg)
	}
}

func broadcastMessage(msg string) {
	mu.Lock()
	defer mu.Unlock()

	for conn := range connections {
		err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			fmt.Println("Error sending message:", err)
			conn.Close()
			delete(connections, conn)
		}
	}
}
