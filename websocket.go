package function

import (
	"log"

	"github.com/gorilla/websocket"
)

// OnOpen is called when a new WebSocket connection is opened.
func OnOpen(conn *websocket.Conn) {
	log.Printf("WebSocket connection opened: %v", conn.RemoteAddr())
}

// OnMessage is called when a new message is received.
func OnMessage(conn *websocket.Conn, messageType int, message []byte) {
	log.Printf("Received message from %v: %s", conn.RemoteAddr(), string(message))

	// Echo the message back to the client
	if err := conn.WriteMessage(messageType, message); err != nil {
		log.Printf("Error sending message: %v", err)
	}
}

// OnClose is called when the WebSocket connection is closed.
func OnClose(conn *websocket.Conn) {
	log.Printf("WebSocket connection closed: %v", conn.RemoteAddr())
	conn.Close()
}

// OnError is called when there is an error with the WebSocket connection.
func OnError(conn *websocket.Conn, err error) {
	log.Printf("WebSocket error from %v: %v", conn.RemoteAddr(), err)
}
