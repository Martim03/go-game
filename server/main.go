package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

const (
	// TODO: Create a server config for values such as update time
	updateClientTimer time.Duration = 1000 * time.Millisecond
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// TODO: implement for production
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// This function is triggered everytime a new connection is made

	// Upgrade the HTTP connection to a WebSocket connection (bidirectional)
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading:", err)
		return
	}
	defer conn.Close()

	go listenLoop(conn)
	go updateLoop(conn)

	for {
	}
}

func listenLoop(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
		}
		fmt.Printf("Received: %s\\n", message)
	}
}

func updateLoop(conn *websocket.Conn) {
	msg := "Tick..."
	for {
		time.Sleep(updateClientTimer)
		if err := conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
			fmt.Println("Error writing message:", err)
		}
	}

}

func startWebSocket() {
	http.HandleFunc("/ws", wsHandler)
	fmt.Println("WebSocket server started on :8080")

	// Start connection on the following port
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func main() {
	startWebSocket()
}
