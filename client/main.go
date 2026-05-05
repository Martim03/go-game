package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
)

const (
	// TODO: Create a client config for values such as update time
	urlScheme = "ws"
	host      = "localhost:8080"
	path      = "/ws"
)

func main() {
	u := url.URL{Scheme: urlScheme, Host: host, Path: path}
	fmt.Printf("Connecting to %s\n", u.String())

	// Connect to the Web Socket
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	defer conn.Close()

	// Send message
	fmt.Println("input: ")
	var ln string
	fmt.Scanf("%s", &ln)
	err = conn.WriteMessage(websocket.TextMessage, []byte(ln))
	if err != nil {
		log.Println("Write error:", err)
		return
	}

	// Listen
	for {
		_, response, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			return
		}
		fmt.Printf("Server echoed: %s\n", response)
	}
}
