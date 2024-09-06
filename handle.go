package function

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true }, // Allow all origins, modify for production
}

// Handle an HTTP Request.
func Handle(w http.ResponseWriter, r *http.Request) {
	// Upgrade the request to a WebSocket connection
	if r.Header.Get("Upgrade") == "websocket" {
		handleWebSocket(w, r)
		return
	}

	// Handle non-WebSocket requests
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Received HTTP request")
	fmt.Printf("%q\n", dump)
	fmt.Fprintf(w, "%q", dump)
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading to WebSocket: %v", err)
		return
	}

	// Call WebSocket handlers
	OnOpen(conn)

	go func() {
		defer OnClose(conn)
		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				OnError(conn, err)
				break
			}
			OnMessage(conn, messageType, message)
		}
	}()
}
